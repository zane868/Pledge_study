// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;
import {Test} from "forge-std/Test.sol";
import {LocalSystem} from "../script/LocalSystem.sol";
contract USDT6Test is Test, LocalSystem {
    function _stableUnit() internal pure override returns(uint256){return 1e6;}
    function _stableArtifact() internal pure override returns(string memory){return "out/TestUSDT.sol/TestUSDT.json";}
    function _start(address sender) internal override {vm.startPrank(sender);}
    function _stop() internal override {vm.stopPrank();}
    function setUp() public {
        operator=makeAddr("operator");ownerA=makeAddr("a");ownerB=makeAddr("b");
        lender=makeAddr("lender");borrower=makeAddr("borrower");fee=makeAddr("fee");
        vm.deal(operator,10000 ether);vm.deal(borrower,10 ether);_deploySystem();
    }
    function fund(uint256 value) internal {
        vm.startPrank(lender);stable.approve(address(pool),value);pool.depositLend(pid,value);vm.stopPrank();
        vm.prank(borrower);pool.depositBorrow{value:1 ether}(pid,1 ether);
        vm.warp(settleTime+1);vm.prank(operator);pool.settle(pid);
    }
    function testSixDecimalMatchingRefundVouchersAndFinalRedemption() public {
        fund(1500e6);
        (uint256 principal,uint256 collateral,,,,)=pool.poolDataInfo(pid);
        assertEq(principal,1000e6);assertEq(collateral,1 ether);
        vm.prank(lender);pool.refundLend(pid);assertEq(stable.balanceOf(lender),19000e6);
        vm.prank(lender);pool.claimLend(pid);assertEq(sp.balanceOf(lender),1000 ether);
        vm.prank(borrower);pool.claimBorrow(pid);assertEq(stable.balanceOf(borrower),1000e6);assertEq(jp.balanceOf(borrower),2000 ether);
        vm.prank(operator);pool.setFee(1000000,2000000);
        vm.warp(endTime+1);vm.prank(operator);pool.finish(pid);
        (,,uint256 due,uint256 remaining,,)=pool.poolDataInfo(pid);assertEq(due,1200e6);
        vm.prank(lender);vm.expectRevert("Zero USDT redemption");pool.withdrawLend(pid,1);
        assertEq(sp.balanceOf(lender),1000 ether);
        assertGe(stable.balanceOf(fee),12e6);
        vm.prank(lender);pool.withdrawLend(pid,250 ether);
        vm.prank(lender);pool.withdrawLend(pid,750 ether);assertEq(stable.balanceOf(lender),20200e6);
        uint256 before=borrower.balance;vm.prank(borrower);pool.withdrawBorrow(pid,2000 ether);
        assertEq(borrower.balance-before,remaining);assertEq(address(pool).balance,0);assertEq(stable.balanceOf(address(pool)),0);
    }
    function testExcessEthAndLateClaimsKeepSixDecimalUnits() public {
        fund(600e6);(,uint256 matched,,,,)=pool.poolDataInfo(pid);assertEq(matched,0.6 ether);
        vm.warp(endTime+1);vm.prank(operator);pool.finish(pid);
        vm.prank(borrower);pool.claimBorrow(pid);assertEq(stable.balanceOf(borrower),600e6);
        uint256 before=borrower.balance;vm.prank(borrower);pool.refundBorrow(pid);assertEq(borrower.balance-before,0.4 ether);
        vm.prank(lender);pool.claimLend(pid);assertEq(sp.balanceOf(lender),600 ether);
        vm.prank(lender);pool.withdrawLend(pid,600 ether);assertEq(stable.balanceOf(lender),20120e6);
    }
    function testRiskUsesUsdtRawUnits() public {
        fund(1000e6);assertFalse(pool.checkoutLiquidate(pid));
        vm.prank(operator);oracle.setPrice(address(0),1100e8);assertTrue(pool.checkoutLiquidate(pid));
        vm.prank(operator);pool.liquidate(pid);(,,,,uint256 due,)=pool.poolDataInfo(pid);assertEq(due,1200e6);
    }
    function testFuzzMatchingDoesNotExceedCollateral(uint32 amount) public {
        uint256 value=bound(uint256(amount),101e6,1000e6);fund(value);
        (uint256 principal,uint256 collateral,,,,)=pool.poolDataInfo(pid);
        assertEq(principal,value);assertLe(collateral,1 ether);
        assertGe(collateral*1000, value*1e12);
    }
}
