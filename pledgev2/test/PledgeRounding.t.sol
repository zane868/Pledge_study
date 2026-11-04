// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {LocalSystem} from "../script/LocalSystem.sol";

/// LocalSystem 的 IPool 未声明这两个 getter，这里补一个最小接口
interface IPoolExtra {
    function userLendInfo(address, uint256) external view returns (uint256, uint256, bool, bool);
    function lendScale(uint256) external view returns (uint256);
}

/**
 * @title 前端镜像公式的取整锁定测试
 * @notice 前端（pledge-fe）为了给出「预计可领 / 预计退款 / 预计赎回」的预览，
 *         在 TS 侧复刻了本合约的乘除顺序。如果前端把两步除法合并成一步，
 *         或者改变取整方向，预览金额就会和链上实际结果不一致 —— 用户看到 A、
 *         实际到账 B，属于会引发纠纷的展示缺陷。
 *
 *         本文件把每条路径的取整规则用精确数字钉死，作为前端实现的对照基准。
 *         前端 shared/src/numbers.ts 中的 expectedSp / expectedJp / expectedLoan /
 *         refundAmount / previewSpRedeem / previewJpRedeem 必须与这里的期望一致。
 *
 * @dev 使用与线上部署一致的 6 位小数 USDT（lendScale = 1e12）。
 *      所有断言用整数运算，避免浮点误差掩盖取整差异。
 */
contract PledgeRoundingTest is Test, LocalSystem {
    uint256 internal constant CAL_DECIMAL = 1e18;
    uint256 internal constant BASE_DECIMAL = 1e8;
    uint256 internal constant MORTGAGE_RATE = 200_000_000; // 200%

    function _stableUnit() internal pure override returns (uint256) {
        return 1e6;
    }

    function _stableArtifact() internal pure override returns (string memory) {
        return "out/TestUSDT.sol/TestUSDT.json";
    }

    function _start(address sender) internal override {
        vm.startPrank(sender);
    }

    function _stop() internal override {
        vm.stopPrank();
    }

    function setUp() public {
        operator = makeAddr("operator");
        ownerA = makeAddr("ownerA");
        ownerB = makeAddr("ownerB");
        lender = makeAddr("lender");
        borrower = makeAddr("borrower");
        recipient = makeAddr("recipient");
        fee = makeAddr("fee");
        vm.deal(operator, 10_000 ether);
        vm.deal(borrower, 1000 ether);
        _deploySystem();
    }

    /// 出借资产是 6 位小数 → lendScale = 10^(18-6) = 1e12
    function _scale() internal pure returns (uint256) {
        return 1e12;
    }

    // ─── 流程助手（与 PledgePool.t.sol / USDT6.t.sol 保持一致）──────────

    function _fund(uint256 lend, uint256 collateral) internal {
        if (lend > 0) {
            vm.startPrank(lender);
            stable.approve(address(pool), lend);
            pool.depositLend(pid, lend);
            vm.stopPrank();
        }
        if (collateral > 0) {
            vm.prank(borrower);
            pool.depositBorrow{value: collateral}(pid, collateral);
        }
    }

    function _settle() internal {
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
    }

    function _claim() internal {
        vm.prank(lender);
        pool.claimLend(pid);
        vm.prank(borrower);
        pool.claimBorrow(pid);
    }

    function _finish() internal {
        vm.warp(endTime + 1);
        vm.prank(operator);
        pool.finish(pid);
    }

    // ─── 1. claimLend 的 SP 铸造量 ─────────────────────────────────────────

    /**
     * 链上：totalSpAmount = settleAmountLend * lendScale
     *       spAmount     = mulDiv(totalSpAmount, stakeAmount, lendSupply)
     *
     * 前端 expectedSp(settleAmountLend, scale, stakeAmount, lendSupply)
     *
     * 关键：出借 1500 USDT、抵押 1 ETH（价值 2000 USDT、抵押率 200%）
     *      → 抵押侧容量 1000 USDT，所以只匹配 1000，出借人拿到 1000 USDT 的 SP。
     */
    function testClaimLendSpAmountMatchesFrontendFormula() public {
        _fund(1500e6, 1 ether);
        _settle();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        assertEq(settleAmountLend, 1000e6);

        uint256 totalSpAmount = settleAmountLend * _scale();

        // 前端公式：mulDiv(settleAmountLend * scale, stakeAmount, lendSupply)
        uint256 expected = (settleAmountLend * _scale()) * 1500e6 / 1500e6;
        assertEq(expected, totalSpAmount);

        vm.prank(lender);
        pool.claimLend(pid);
        assertEq(sp.balanceOf(lender), expected);
        assertEq(sp.balanceOf(lender), 1000 ether);
    }

    /**
     * 取整方向必须向下：多人参与时每人拿到的 SP 之和可能小于总 SP，
     * 差额留在合约里。前端预览用同样的向下取整语义。
     */
    function testClaimLendRoundsDownAndLeavesRemainder() public {
        address second = makeAddr("secondLender");
        address third = makeAddr("thirdLender");

        uint256[3] memory stakes = [uint256(100000001), uint256(500000001), uint256(900000001)];
        address[3] memory lenders = [lender, second, third];

        // 给另外两个出借人发 USDT
        vm.startPrank(operator);
        stable.transfer(second, stakes[1]);
        stable.transfer(third, stakes[2]);
        vm.stopPrank();

        for (uint256 i; i < 3; i++) {
            vm.startPrank(lenders[i]);
            stable.approve(address(pool), stakes[i]);
            pool.depositLend(pid, stakes[i]);
            vm.stopPrank();
        }

        vm.prank(borrower);
        pool.depositBorrow{value: 1 ether}(pid, 1 ether);
        _settle();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        uint256 totalSpAmount = settleAmountLend * _scale();
        uint256 totalStaked = 1500000003;

        uint256 sum;
        for (uint256 i; i < 3; i++) {
            // 前端公式：mulDiv(totalSpAmount, stakeAmount, lendSupply)
            uint256 expected = totalSpAmount * stakes[i] / totalStaked;
            vm.prank(lenders[i]);
            pool.claimLend(pid);
            assertEq(sp.balanceOf(lenders[i]), expected);
            sum += expected;
        }
        // 向下取整导致分给用户的总额小于总 SP，差额留在合约里
        assertLe(sum, totalSpAmount);
        assertLt(sum, totalSpAmount);
    }

    // ─── 2. claimBorrow 的 JP 铸造量与借款支付额 ────────────────────────

    /**
     * 链上：
     *   totalJpAmount = mulDiv(settleAmountLend * lendScale, martgageRate, baseDecimal)
     *   jpAmount      = mulDiv(totalJpAmount, stakeAmount, borrowSupply)
     *   borrowAmount  = mulDiv(settleAmountLend, stakeAmount, borrowSupply)
     *
     * 前端 expectedJp / expectedLoan
     *
     * 关键：JP 凭证是 18 位，借款是 6 位，两者的分母不同，
     *      前端不能用借款额反推 JP 数量。
     */
    function testClaimBorrowJpAndLoanMatchFrontendFormula() public {
        _fund(1500e6, 1 ether);
        _settle();

        (uint256 settleAmountLend, uint256 settleAmountBorrow,,,,) = pool.poolDataInfo(pid);
        assertEq(settleAmountLend, 1000e6);
        assertEq(settleAmountBorrow, 1 ether);

        // 前端 jpDenominator: mulDiv(settleAmountLend * scale, martgageRate, baseDecimal)
        uint256 totalJpAmount = (settleAmountLend * _scale()) * MORTGAGE_RATE / BASE_DECIMAL;
        // 前端 expectedJp: mulDiv(totalJpAmount, stakeAmount, borrowSupply)
        uint256 expectedJp = totalJpAmount * 1 ether / 1 ether;
        // 前端 expectedLoan: mulDiv(settleAmountLend, stakeAmount, borrowSupply)
        uint256 expectedLoan = settleAmountLend * 1 ether / 1 ether;

        vm.prank(borrower);
        pool.claimBorrow(pid);

        assertEq(jp.balanceOf(borrower), expectedJp);
        assertEq(stable.balanceOf(borrower), expectedLoan);
        assertEq(jp.balanceOf(borrower), 2000 ether);
        assertEq(stable.balanceOf(borrower), 1000e6);
    }

    /**
     * 多借款人场景下，JP 与借款金额各自独立向下取整，两者不可互相推导。
     * 前端必须分别计算 expectedJp 和 expectedLoan。
     */
    function testClaimBorrowRoundsJpAndLoanIndependently() public {
        address second = makeAddr("secondBorrower");
        vm.deal(second, 10 ether);

        vm.startPrank(lender);
        stable.approve(address(pool), 1000e6);
        pool.depositLend(pid, 1000e6);
        vm.stopPrank();

        vm.prank(borrower);
        pool.depositBorrow{value: 1 ether}(pid, 1 ether);
        vm.prank(second);
        pool.depositBorrow{value: 0.5 ether}(pid, 0.5 ether);

        _settle();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        uint256 totalJpAmount = (settleAmountLend * _scale()) * MORTGAGE_RATE / BASE_DECIMAL;
        uint256 borrowSupply = 1.5 ether;

        // 第一个借款人的份额
        uint256 jpExpected = totalJpAmount * 1 ether / borrowSupply;
        uint256 loanExpected = settleAmountLend * 1 ether / borrowSupply;

        vm.prank(borrower);
        pool.claimBorrow(pid);
        assertEq(jp.balanceOf(borrower), jpExpected);
        assertEq(stable.balanceOf(borrower), loanExpected);

        // 第二个借款人
        uint256 jpExpected2 = totalJpAmount * 0.5 ether / borrowSupply;
        uint256 loanExpected2 = settleAmountLend * 0.5 ether / borrowSupply;
        vm.prank(second);
        pool.claimBorrow(pid);
        assertEq(jp.balanceOf(second), jpExpected2);
        assertEq(stable.balanceOf(second), loanExpected2);

        // 两人之和不超过总额（向下取整）
        assertLe(jpExpected + jpExpected2, totalJpAmount);
        assertLe(loanExpected + loanExpected2, settleAmountLend);
    }

    // ─── 3. refundLend / refundBorrow 的两步取整 ─────────────────────────

    /**
     * 链上：
     *   userShare    = floor(stakeAmount * calDecimal / lendSupply)
     *   refundAmount = floor((lendSupply - settleAmountLend) * userShare / calDecimal)
     *
     * 关键：两步都向下取整，不能合并成一次除法。
     * 前端 refundAmount(stakeAmount, lendSupply, settleAmountLend)
     *
     * 1500 USDT 出借 / 1 ETH 抵押 → 匹配 1000，未成交 500。
     */
    function testRefundLendPartialUnmatchedAmount() public {
        _fund(1500e6, 1 ether);
        _settle();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        assertEq(settleAmountLend, 1000e6);

        uint256 lendSupply = 1500e6;
        uint256 userShare = 1500e6 * CAL_DECIMAL / lendSupply;
        uint256 expectedRefund = (lendSupply - settleAmountLend) * userShare / CAL_DECIMAL;
        assertEq(expectedRefund, 500e6);

        uint256 before = stable.balanceOf(lender);
        vm.prank(lender);
        pool.refundLend(pid);
        assertEq(stable.balanceOf(lender) - before, expectedRefund);

        // 链上记录的 hasNoRefund 与预览一致
        (,, bool noRefund,) = IPoolExtra(address(pool)).userLendInfo(lender, pid);
        assertTrue(noRefund);
    }

    /**
     * 两人出借时，退款按各自本金占比分配，且每一步都向下取整。
     * 这里锁住「两步取整」与「合并取整」的区别 —— 前端若合并取整会差最小单位。
     */
    function testRefundLendTwoStepRoundingAcrossMultipleLenders() public {
        address second = makeAddr("secondLender");
        vm.startPrank(operator);
        stable.transfer(second, 700e6);
        vm.stopPrank();

        vm.startPrank(lender);
        stable.approve(address(pool), 1500e6);
        pool.depositLend(pid, 1500e6);
        vm.stopPrank();

        vm.startPrank(second);
        stable.approve(address(pool), 700e6);
        pool.depositLend(pid, 700e6);
        vm.stopPrank();

        vm.prank(borrower);
        pool.depositBorrow{value: 1 ether}(pid, 1 ether);
        _settle();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        assertEq(settleAmountLend, 1000e6);

        uint256 lendSupply = 2200e6;
        uint256 unmatched = lendSupply - settleAmountLend;

        // 前端 refundAmount(stake, lendSupply, settleAmountLend)
        uint256 shareA = 1500e6 * CAL_DECIMAL / lendSupply;
        uint256 refundA = unmatched * shareA / CAL_DECIMAL;
        uint256 shareB = 700e6 * CAL_DECIMAL / lendSupply;
        uint256 refundB = unmatched * shareB / CAL_DECIMAL;

        uint256 beforeA = stable.balanceOf(lender);
        vm.prank(lender);
        pool.refundLend(pid);
        assertEq(stable.balanceOf(lender) - beforeA, refundA);

        uint256 beforeB = stable.balanceOf(second);
        vm.prank(second);
        pool.refundLend(pid);
        assertEq(stable.balanceOf(second) - beforeB, refundB);

        // 两人拿到的退款之和不超过未成交总额
        assertLe(refundA + refundB, unmatched);
    }

    /**
     * refundBorrow 的分母是 borrowSupply，不是凭证总供应量。
     * 前端 refundAmount(borrowStake, pool.borrowSupply, settleAmountBorrow)
     */
    function testRefundBorrowUsesBorrowSupplyDenominator() public {
        _fund(600e6, 1 ether);
        _settle();

        (, uint256 settleAmountBorrow,,,,) = pool.poolDataInfo(pid);
        assertEq(settleAmountBorrow, 0.6 ether);

        uint256 borrowSupply = 1 ether;
        uint256 userShare = 1 ether * CAL_DECIMAL / borrowSupply;
        uint256 expectedRefund = (borrowSupply - settleAmountBorrow) * userShare / CAL_DECIMAL;
        assertEq(expectedRefund, 0.4 ether);

        uint256 before = borrower.balance;
        vm.prank(borrower);
        pool.refundBorrow(pid);
        assertEq(borrower.balance - before, expectedRefund);
    }

    // ─── 4. 赎回：分母不同，且 SP 侧有 require(> 0) ────────────────────

    /**
     * withdrawLend 的分母是 totalSpAmount = settleAmountLend * lendScale，
     * withdrawBorrow 的分母是 totalJpAmount。两者不同，不能混用。
     *
     * 前端 previewSpRedeem / previewJpRedeem
     */
    function testRedeemDenominatorsDifferBetweenSpAndJp() public {
        _fund(1000e6, 1 ether);
        _settle();
        _claim();

        (uint256 settleAmountLend,,,,,) = pool.poolDataInfo(pid);
        uint256 totalSpAmount = settleAmountLend * _scale();
        uint256 totalJpAmount = (settleAmountLend * _scale()) * MORTGAGE_RATE / BASE_DECIMAL;

        // 两者必须不同，否则前端一旦混用分母也不会被发现
        assertNotEq(totalSpAmount, totalJpAmount);
        assertEq(totalJpAmount, totalSpAmount * 2);

        _finish();

        (,, uint256 finishAmountLend, uint256 finishAmountBorrow,,) = pool.poolDataInfo(pid);
        assertGt(finishAmountLend, 0);

        // 前端 previewSpRedeem: mulDiv(finishAmountLend, spAmount, totalSpAmount)
        uint256 spExpected = finishAmountLend * 1000 ether / totalSpAmount;
        uint256 before = stable.balanceOf(lender);
        vm.prank(lender);
        pool.withdrawLend(pid, 1000 ether);
        assertEq(stable.balanceOf(lender) - before, spExpected);

        // 前端 previewJpRedeem: mulDiv(finishAmountBorrow, jpAmount, totalJpAmount)
        uint256 jpExpected = finishAmountBorrow * 2000 ether / totalJpAmount;
        before = borrower.balance;
        vm.prank(borrower);
        pool.withdrawBorrow(pid, 2000 ether);
        assertEq(borrower.balance - before, jpExpected);
    }

    /**
     * SP 侧有 require(redeemAmount > 0)，零兑付必须回滚 —— 前端据此禁用按钮。
     * JP 侧没有该检查，零兑付会成功执行并销毁凭证 —— 属 DEC-14，前端必须披露。
     */
    function testZeroRedemptionRevertsOnSpButSucceedsOnJp() public {
        _fund(1000e6, 1 ether);
        _settle();
        _claim();
        _finish();

        // SP：数量大于 0 但按比例算出的兑付额为 0
        vm.prank(lender);
        vm.expectRevert("Zero USDT redemption");
        pool.withdrawLend(pid, 1);
        assertEq(sp.balanceOf(lender), 1000 ether); // 回滚，凭证未被销毁

        // JP：同样的最小数量可以成功执行（兑付额为 0，凭证被销毁）
        uint256 jpBefore = jp.balanceOf(borrower);
        assertGt(jpBefore, 0);
        vm.prank(borrower);
        pool.withdrawBorrow(pid, 1);
        assertEq(jp.balanceOf(borrower), jpBefore - 1);
    }

    // ─── 5. lendScale 是链上事实，不是前端常量 ──────────────────────────

    /**
     * lendScale = 10^(18 - 出借资产小数位)。前端此前把它当作部署常量，
     * 风险是池的出借资产小数位变化后预览全部错位。
     * 6 位 USDT → 1e12；链上 getter 是权威来源。
     */
    function testLendScaleIsDerivedFromLendingAssetDecimals() public view {
        assertEq(IPoolExtra(address(pool)).lendScale(pid), _scale());
        assertEq(IPoolExtra(address(pool)).lendScale(pid), 1e12);
    }

    // ─── 6. 紧急取回：原额退回，标记为已退款 ────────────────────────────

    /**
     * emergencyLendWithdrawal / emergencyBorrowWithdrawal 直接退回本金，
     * 并把 hasNoRefund 置为 true —— 前端按钮据此禁用。
     */
    function testEmergencyWithdrawalReturnsPrincipalAndBlocksRefund() public {
        _fund(1000e6, 0);
        _settle();
        assertEq(pool.getPoolState(pid), 4);

        uint256 before = stable.balanceOf(lender);
        vm.prank(lender);
        pool.emergencyLendWithdrawal(pid);
        assertEq(stable.balanceOf(lender) - before, 1000e6);

        (uint256 stakeAmount, uint256 refundAmount, bool noRefund,) =
            IPoolExtra(address(pool)).userLendInfo(lender, pid);
        assertEq(stakeAmount, 1000e6);
        assertTrue(noRefund);
        // 紧急取回直接返还本金，不累加 refundAmount —— 前端不能把它当作「已领退款」显示
        assertEq(refundAmount, 0);

        // 不能再走普通退款：UNDONE 状态下 refundLend 被状态修饰符挡住，
        // 且 hasNoRefund 已置位 —— 两条防线都在
        vm.prank(lender);
        vm.expectRevert("state: not match and undone");
        pool.refundLend(pid);

        // 紧急取回同样不能再执行一次
        vm.prank(lender);
        vm.expectRevert("refundLend: again refund");
        pool.emergencyLendWithdrawal(pid);
    }
}
