// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {LocalSystem} from "../script/LocalSystem.sol";

/**
 * @dev 覆盖 forceSettle / forceFinish：这两个接口只跳过时间校验，
 *      其余状态机、计价、交换逻辑必须与 settle / finish 完全一致。
 */
contract ForceSettlementTest is Test, LocalSystem {
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
        vm.deal(borrower, 100 ether);
        _deploySystem();
    }

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

    /// @dev 常规 settle 在时间未到时必须失败
    function testSettleRevertsBeforeSettleTime() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime - 1);
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"settle: 小于结算时间"));
        pool.settle(pid);
    }

    /// @dev 强制结算：时间未到也能结算，状态推进到 EXECUTION
    function testForceSettleBeforeSettleTime() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime - 1);
        vm.prank(operator);
        pool.forceSettle(pid);
        assertEq(pool.getPoolState(pid), 1); // EXECUTION
    }

    /// @dev 强制结算仍然必须通过多签校验
    function testForceSettleStillRequiresMultisig() public {
        _fund(1000 ether, 10 ether);
        vm.prank(borrower);
        vm.expectRevert(bytes("multiSignatureClient : This tx is not aprroved"));
        pool.forceSettle(pid);
    }

    /// @dev 强制结算仍然要求池子处于 MATCH 状态
    function testForceSettleRejectsNonMatchState() public {
        _fund(1000 ether, 10 ether);
        vm.prank(operator);
        pool.forceSettle(pid);
        assertEq(pool.getPoolState(pid), 1);
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"settle: 池子状态必须是匹配"));
        pool.forceSettle(pid);
    }

    /// @dev 常规 finish 在时间未到时必须失败
    function testFinishRevertsBeforeEndTime() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
        vm.warp(endTime - 1);
        vm.prank(operator);
        vm.expectRevert(bytes("finish: less than end time"));
        pool.finish(pid);
    }

    /// @dev 强制结算：时间未到也能结束池子
    function testForceFinishBeforeEndTime() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
        vm.warp(endTime - 1);
        vm.prank(operator);
        pool.forceFinish(pid);
        assertEq(pool.getPoolState(pid), 2); // FINISH
    }

    /// @dev 强制结算仍然必须通过多签校验
    function testForceFinishStillRequiresMultisig() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
        vm.prank(borrower);
        vm.expectRevert(bytes("multiSignatureClient : This tx is not aprroved"));
        pool.forceFinish(pid);
    }

    /// @dev 强制结算仍然要求池子处于 EXECUTION 状态
    function testForceFinishRejectsNonExecutionState() public {
        _fund(1000 ether, 10 ether);
        vm.prank(operator);
        vm.expectRevert(bytes("finish: pool state must be execution"));
        pool.forceFinish(pid);
    }

    /// @dev 强制结算的计息与常规 finish 一致（同样按 endTime - settleTime 计息）
    function testForceFinishInterestMatchesRegularFinish() public {
        _fund(1000 ether, 10 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
        vm.warp(endTime - 1);
        vm.prank(operator);
        pool.forceFinish(pid);

        (uint256 settleAmountLend,, uint256 finishAmountLend,,,) = pool.poolDataInfo(pid);
        // 利息非零即可，具体数值由时间比率决定
        assertGt(finishAmountLend, settleAmountLend);
    }

    // ---------------------------------------------------------------
    // forceLiquidate
    // ---------------------------------------------------------------

    /// @dev 强制结算会把 settleTime 提前到当前时间前 5 秒，
    ///      因此普通 liquidate 的时间校验（block.timestamp > settleTime）也能通过，
    ///      不再被时间闸门拦截。
    function testForceSettleAdvancesSettleTimeSoLiquidatePasses() public {
        _fund(1000 ether, 1 ether);
        vm.warp(settleTime - 1);
        vm.prank(operator);
        pool.forceSettle(pid);
        assertEq(pool.getPoolState(pid), 1); // EXECUTION

        vm.prank(operator);
        pool.liquidate(pid);
        assertEq(pool.getPoolState(pid), 3); // LIQUIDATION
    }

    /// @dev 强制清算：结算时间之前也能清算
    function testForceLiquidateBeforeSettleTime() public {
        _fund(1000 ether, 1 ether);
        vm.warp(settleTime - 1);
        vm.prank(operator);
        pool.forceSettle(pid);
        assertEq(pool.getPoolState(pid), 1); // EXECUTION

        vm.prank(operator);
        pool.forceLiquidate(pid);
        assertEq(pool.getPoolState(pid), 3); // LIQUIDATION
    }

    /// @dev 强制清算仍然必须通过多签校验
    function testForceLiquidateStillRequiresMultisig() public {
        _fund(1000 ether, 1 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);
        vm.prank(borrower);
        vm.expectRevert(bytes("multiSignatureClient : This tx is not aprroved"));
        pool.forceLiquidate(pid);
    }

    /// @dev 强制清算仍然要求池子处于 EXECUTION 状态
    function testForceLiquidateRejectsNonExecutionState() public {
        _fund(1000 ether, 1 ether);
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"liquidate: 池子的状态必须是执行状态"));
        pool.forceLiquidate(pid);
    }

    /// @dev 强制清算与常规清算的分配结果一致
    function testForceLiquidateMatchesRegularLiquidate() public {
        _fund(1000 ether, 1 ether);
        vm.warp(settleTime + 1);
        vm.prank(operator);
        pool.settle(pid);

        vm.prank(operator);
        pool.forceLiquidate(pid);
        assertEq(pool.getPoolState(pid), 3); // LIQUIDATION
        (,,,, uint256 allocated, uint256 remaining) = pool.poolDataInfo(pid);
        assertEq(allocated, 1200 ether);
        assertGt(remaining, 0);
    }
}
