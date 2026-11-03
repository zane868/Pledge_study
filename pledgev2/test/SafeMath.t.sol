// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {SafeMath} from "../contracts/library/SafeMath.sol";

contract MathHarness {
    function add(uint256 a, uint256 b) external pure returns (uint256) {
        return SafeMath.add(a, b);
    }

    function mul(uint256 a, uint256 b) external pure returns (uint256) {
        return SafeMath.mul(a, b);
    }

    function sub(uint256 a, uint256 b) external pure returns (uint256) {
        return SafeMath.sub(a, b);
    }
}

contract SafeMathUpgradeTest is Test {
    MathHarness internal math = new MathHarness();

    function testOriginalOverflowAndUnderflowRevertMessages() public {
        vm.expectRevert("SafeMath: addition overflow");
        math.add(type(uint256).max, 1);
        vm.expectRevert("SafeMath: multiplication overflow");
        math.mul(type(uint256).max, 2);
        vm.expectRevert("SafeMath: subtraction overflow");
        math.sub(0, 1);
    }

    function testFuzzGuardedArithmeticMatchesNative(uint128 a, uint128 b) public view {
        assertEq(math.add(a, b), uint256(a) + uint256(b));
        assertEq(math.mul(a, b), uint256(a) * uint256(b));
        if (a >= b) assertEq(math.sub(a, b), uint256(a) - uint256(b));
    }
}
