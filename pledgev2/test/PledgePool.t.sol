// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {stdStorage, StdStorage} from "forge-std/StdStorage.sol";
import {LocalSystem} from "../script/LocalSystem.sol";

interface IUpgradeFactory {
    function getPair(address, address) external view returns (address);
}

interface IUpgradePair {
    function getReserves() external view returns (uint112, uint112, uint32);
    function price0CumulativeLast() external view returns (uint256);
    function price1CumulativeLast() external view returns (uint256);
    function sync() external;
}

contract PledgePoolTest is Test, LocalSystem {
    using stdStorage for StdStorage;
    string internal constant UNAPPROVED = "multiSignatureClient : This tx is not aprroved";

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

    function testDeploymentAndPoolMinters() public view {
        assertEq(pool.getMultiSignatureAddress(), address(multisig));
        assertEq(pool.oracle(), address(oracle));
        assertEq(pool.poolLength(), 1);
        assertTrue(sp.isMinter(address(pool)));
        assertTrue(jp.isMinter(address(pool)));
        assertFalse(sp.isMinter(operator));
    }

    function testMultisigThresholdAndRevocation() public {
        vm.startPrank(borrower);
        vm.expectRevert(bytes(UNAPPROVED));
        pool.setFee(1, 2);
        multisig.createApplication(address(pool));
        vm.stopPrank();
        bytes32 hash = multisig.getApplicationHash(borrower, address(pool));
        vm.prank(ownerA);
        multisig.signApplication(hash);
        vm.prank(borrower);
        vm.expectRevert(bytes(UNAPPROVED));
        pool.setFee(1, 2);
        vm.prank(ownerB);
        multisig.signApplication(hash);
        vm.prank(borrower);
        pool.setFee(1, 2);
        assertEq(pool.borrowFee(), 2);
        vm.prank(ownerA);
        multisig.revokeSignApplication(hash);
        vm.prank(borrower);
        vm.expectRevert(bytes(UNAPPROVED));
        pool.setFee(3, 4);
    }

    function testUnauthorizedPriceMintBurn() public {
        vm.startPrank(borrower);
        vm.expectRevert(bytes(UNAPPROVED));
        oracle.setPrice(address(0), 1);
        vm.expectRevert(bytes(UNAPPROVED));
        sp.addMinter(borrower);
        vm.expectRevert("Token: caller is not the minter");
        sp.mint(borrower, 1);
        vm.expectRevert("Token: caller is not the minter");
        jp.burn(borrower, 1);
        vm.stopPrank();
    }

    function testMinimumCapAndStrictTimeBoundaries() public {
        vm.startPrank(lender);
        stable.approve(address(pool), 20_000 ether);
        vm.expectRevert(bytes(unicode"depositLend: 少于最小金额"));
        pool.depositLend(pid, 100 ether);
        vm.expectRevert(bytes(unicode"depositLend: 数量超过限制"));
        pool.depositLend(pid, 100_001 ether);
        vm.stopPrank();
        _fund(1000 ether, 1 ether);
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"settle: 小于结算时间"));
        pool.settle(pid);
        vm.warp(settleTime);
        vm.prank(borrower);
        vm.expectRevert("Less than this time");
        pool.depositBorrow{value: 1}(pid, 1);
        assertFalse(pool.checkoutSettle(pid));
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"settle: 小于结算时间"));
        pool.settle(pid);
        _settle();
        assertEq(pool.getPoolState(pid), 1);
        vm.prank(operator);
        vm.expectRevert(bytes(unicode"settle: 池子状态必须是匹配"));
        pool.settle(pid);
    }

    function testExcessStableRefundAndLoanClaims() public {
        _fund(1500 ether, 1 ether);
        _settle();
        (uint256 matched, uint256 collateral,,,,) = pool.poolDataInfo(pid);
        assertEq(matched, 1000 ether);
        assertEq(collateral, 1 ether);
        uint256 before = stable.balanceOf(lender);
        vm.startPrank(lender);
        pool.refundLend(pid);
        assertEq(stable.balanceOf(lender) - before, 500 ether);
        vm.expectRevert("refundLend: repeat refund");
        pool.refundLend(pid);
        pool.claimLend(pid);
        vm.expectRevert(bytes(unicode"claimLend: 不能再次领取"));
        pool.claimLend(pid);
        vm.stopPrank();
        vm.startPrank(borrower);
        pool.claimBorrow(pid);
        vm.expectRevert(bytes(unicode"claimBorrow: 再次索取"));
        pool.claimBorrow(pid);
        vm.stopPrank();
        assertEq(stable.balanceOf(borrower), 1000 ether);
        assertEq(sp.balanceOf(lender), 1000 ether);
        assertEq(jp.balanceOf(borrower), 2000 ether);
    }

    function testUnmatchedEthRefund() public {
        _fund(600 ether, 1 ether);
        _settle();
        (, uint256 collateral,,,,) = pool.poolDataInfo(pid);
        assertEq(collateral, 0.6 ether);
        uint256 before = borrower.balance;
        vm.startPrank(borrower);
        pool.refundBorrow(pid);
        assertEq(borrower.balance - before, 0.4 ether);
        pool.claimBorrow(pid);
        vm.expectRevert("refundBorrow: again refund");
        pool.refundBorrow(pid);
        vm.stopPrank();
        assertEq(stable.balanceOf(borrower), 600 ether);
    }

    function testMultipleLendersAndBorrowers() public {
        _fund(1000 ether, 0.5 ether);
        vm.prank(operator);
        stable.transfer(recipient, 1000 ether);
        vm.startPrank(recipient);
        stable.approve(address(pool), 1000 ether);
        pool.depositLend(pid, 1000 ether);
        vm.stopPrank();
        vm.prank(operator);
        pool.depositBorrow{value: 0.5 ether}(pid, 0.5 ether);
        _settle();
        address[2] memory lenders = [lender, recipient];
        address[2] memory borrowers = [borrower, operator];
        for (uint256 i; i < 2; i++) {
            uint256 before = stable.balanceOf(lenders[i]);
            vm.startPrank(lenders[i]);
            pool.refundLend(pid);
            pool.claimLend(pid);
            vm.stopPrank();
            assertEq(stable.balanceOf(lenders[i]) - before, 500 ether);
            assertEq(sp.balanceOf(lenders[i]), 500 ether);
            before = stable.balanceOf(borrowers[i]);
            vm.prank(borrowers[i]);
            pool.claimBorrow(pid);
            assertEq(stable.balanceOf(borrowers[i]) - before, 500 ether);
            assertEq(jp.balanceOf(borrowers[i]), 1000 ether);
        }
        _finish();
        (,,, uint256 remaining,,) = pool.poolDataInfo(pid);
        for (uint256 i; i < 2; i++) {
            uint256 before = stable.balanceOf(lenders[i]);
            vm.prank(lenders[i]);
            pool.withdrawLend(pid, 500 ether);
            assertEq(stable.balanceOf(lenders[i]) - before, 600 ether);
            before = borrowers[i].balance;
            vm.prank(borrowers[i]);
            pool.withdrawBorrow(pid, 1000 ether);
            assertEq(borrowers[i].balance - before, remaining / 2);
        }
        assertEq(stable.balanceOf(address(pool)), 0);
        assertEq(address(pool).balance, remaining % 2);
    }

    function testMintRevertRollsBackLoanClaim() public {
        _fund(1000 ether, 1 ether);
        _settle();
        vm.prank(operator);
        jp.delMinter(address(pool));
        vm.prank(borrower);
        vm.expectRevert("Token: caller is not the minter");
        pool.claimBorrow(pid);
        (,,, bool claimed) = pool.userBorrowInfo(borrower, pid);
        assertFalse(claimed);
        assertEq(stable.balanceOf(borrower), 0);
        assertEq(jp.totalSupply(), 0);
        vm.prank(operator);
        jp.addMinter(address(pool));
        vm.prank(borrower);
        pool.claimBorrow(pid);
        assertEq(stable.balanceOf(borrower), 1000 ether);
    }

    function testThirtyDayIntegerInterest() public {
        _createPool(30 days);
        _fund(1000 ether, 1 ether);
        _settle();
        _finish();
        (,, uint256 allocated,,,) = pool.poolDataInfo(pid);
        uint256 timeRatio = uint256(30 days) * 1e8 / YEAR;
        assertEq(allocated, 1000 ether + timeRatio * 20_000_000 * 1000 ether / 1e16);
        (uint256 original,,,,,) = pool.poolDataInfo(0);
        assertEq(original, 0);
    }

    function testFundraisingLenderExitOnFailure() public {
        _fund(1000 ether, 0);
        _settle();
        assertEq(pool.getPoolState(pid), 4);
        uint256 before = stable.balanceOf(lender);
        vm.startPrank(lender);
        pool.emergencyLendWithdrawal(pid);
        assertEq(stable.balanceOf(lender) - before, 1000 ether);
        vm.expectRevert("refundLend: again refund");
        pool.emergencyLendWithdrawal(pid);
        vm.stopPrank();
    }

    function testFundraisingBorrowerExitOnFailure() public {
        _fund(0, 1 ether);
        _settle();
        assertEq(pool.getPoolState(pid), 4);
        uint256 before = borrower.balance;
        vm.startPrank(borrower);
        pool.emergencyBorrowWithdrawal(pid);
        assertEq(borrower.balance - before, 1 ether);
        vm.expectRevert("refundBorrow: again refund");
        pool.emergencyBorrowWithdrawal(pid);
        vm.stopPrank();
    }

    function testDexFeesTransferAndPartialRedemption() public {
        _fund(1000 ether, 1 ether);
        _settle();
        _claim();
        vm.prank(lender);
        sp.transfer(recipient, 250 ether);
        vm.prank(operator);
        pool.setFee(1_000_000, 2_000_000);
        vm.warp(endTime);
        vm.prank(operator);
        vm.expectRevert("finish: less than end time");
        pool.finish(pid);
        address[] memory path = new address[](2);
        path[0] = weth;
        path[1] = address(stable);
        uint256 sold = router.getAmountsIn(1212 ether, path)[0];
        uint256 received = router.getAmountsOut(sold, path)[1];
        uint256 remaining = 1 ether - sold;
        uint256 ethFee = remaining * 2_000_000 / 1e8;
        _finish();
        (,, uint256 allocated, uint256 ethAllocated,,) = pool.poolDataInfo(pid);
        assertEq(pool.getPoolState(pid), 2);
        assertEq(allocated, 1200 ether);
        assertEq(ethAllocated, remaining - ethFee);
        assertEq(stable.balanceOf(fee), received - 1200 ether);
        assertEq(fee.balance, ethFee);
        _redeemSp(recipient, 100 ether, 120 ether);
        _redeemSp(lender, 750 ether, 900 ether);
        _redeemSp(recipient, 150 ether, 180 ether);
        uint256 before = borrower.balance;
        vm.prank(borrower);
        pool.withdrawBorrow(pid, 2000 ether);
        assertEq(borrower.balance - before, ethAllocated);
        assertEq(sp.totalSupply(), 0);
        assertEq(jp.totalSupply(), 0);
        assertEq(stable.balanceOf(address(pool)), 0);
        assertEq(address(pool).balance, 0);
        vm.prank(operator);
        vm.expectRevert("finish: pool state must be execution");
        pool.finish(pid);
    }

    function _redeemSp(address holder, uint256 amount, uint256 payment) internal {
        uint256 before = stable.balanceOf(holder);
        vm.prank(holder);
        pool.withdrawLend(pid, amount);
        assertEq(stable.balanceOf(holder) - before, payment);
    }

    function testLateClaimsAfterFinish() public {
        _fund(600 ether, 1 ether);
        _settle();
        _finish();
        _claim();
        assertEq(stable.balanceOf(borrower), 600 ether);
        uint256 before = borrower.balance;
        vm.prank(borrower);
        pool.refundBorrow(pid);
        assertEq(borrower.balance - before, 0.4 ether);
        _redeemSp(lender, 600 ether, 720 ether);
        (,,, uint256 remaining,,) = pool.poolDataInfo(pid);
        before = borrower.balance;
        vm.prank(borrower);
        pool.withdrawBorrow(pid, 1200 ether);
        assertEq(borrower.balance - before, remaining);
        assertEq(stable.balanceOf(address(pool)), 0);
        assertEq(address(pool).balance, 0);
    }

    function testPauseAndAuthorizedSettlement() public {
        _fund(1000 ether, 1 ether);
        vm.prank(operator);
        pool.setPause();
        vm.prank(borrower);
        vm.expectRevert("Stake has been suspended");
        pool.depositBorrow{value: 1}(pid, 1);
        _settle();
        vm.prank(lender);
        vm.expectRevert("Stake has been suspended");
        pool.claimLend(pid);
        vm.prank(operator);
        pool.setPause();
        vm.prank(lender);
        pool.claimLend(pid);
        assertEq(sp.balanceOf(lender), 1000 ether);
    }

    function testZeroPriceRejected() public {
        _fund(1000 ether, 1 ether);
        vm.prank(operator);
        vm.expectRevert("Invalid price");
        oracle.setPrice(address(stable), 0);
        assertEq(pool.getPoolState(pid), 0);
    }

    function testExistingLiquidationThresholdAndFullTermInterest() public {
        _fund(1000 ether, 1 ether);
        _settle();
        assertFalse(pool.checkoutLiquidate(pid));
        vm.startPrank(operator);
        oracle.setPrice(address(0), 1100e8);
        assertTrue(pool.checkoutLiquidate(pid));
        oracle.setPrice(address(0), 2000e8);
        vm.stopPrank();
        _claim();
        // Source behavior: threshold is advisory; this documents a PRD gap.
        vm.prank(operator);
        pool.liquidate(pid);
        assertEq(pool.getPoolState(pid), 3);
        (,,,, uint256 allocated, uint256 remaining) = pool.poolDataInfo(pid);
        assertEq(allocated, 1200 ether);
        _redeemSp(lender, 1000 ether, 1200 ether);
        uint256 before = borrower.balance;
        vm.prank(borrower);
        pool.withdrawBorrow(pid, 2000 ether);
        assertEq(borrower.balance - before, remaining);
    }

    function testDexFailureRollsBackFinish() public {
        _fund(1000 ether, 1 ether);
        _settle();
        vm.prank(operator);
        pool.setSwapRouterAddress(recipient);
        uint256 before = address(pool).balance;
        vm.warp(endTime + 1);
        vm.prank(operator);
        vm.expectRevert();
        pool.finish(pid);
        assertEq(pool.getPoolState(pid), 1);
        (,, uint256 allocated,,,) = pool.poolDataInfo(pid);
        assertEq(allocated, 0);
        assertEq(address(pool).balance, before);
    }

    function testDexTimestampAndPriceAccumulatorWrapping() public {
        IUpgradePair pair = IUpgradePair(IUpgradeFactory(factory).getPair(weth, address(stable)));
        vm.warp(uint256(type(uint32).max) - 1);
        pair.sync();
        (uint112 reserve0, uint112 reserve1, uint32 previousTime) = pair.getReserves();
        assertEq(previousTime, type(uint32).max - 1);
        uint256 initial = type(uint256).max - 3;
        stdstore.target(address(pair)).sig("price0CumulativeLast()").checked_write(initial);
        stdstore.target(address(pair)).sig("price1CumulativeLast()").checked_write(initial);
        vm.warp(uint256(type(uint32).max) + 6);
        pair.sync();
        (,, uint32 currentTime) = pair.getReserves();
        assertEq(currentTime, 5);
        uint256 expected0;
        uint256 expected1;
        unchecked {
            expected0 = initial + ((uint256(reserve1) << 112) / reserve0) * 7;
            expected1 = initial + ((uint256(reserve0) << 112) / reserve1) * 7;
        }
        assertEq(pair.price0CumulativeLast(), expected0);
        assertEq(pair.price1CumulativeLast(), expected1);
    }
}
