// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../contracts/mock/TestUSDT.sol";

contract DeployNewTestUSDT is Script {
    function run() external {
        vm.startBroadcast();

        TestUSDT stable = new TestUSDT();

        vm.stopBroadcast();

        console.log("New TestUSDT deployed at:", address(stable));
        console.log("Name:", stable.name());
        console.log("Symbol:", stable.symbol());
        console.log("Decimals:", stable.decimals());
        console.log("Total Supply:", stable.totalSupply());
        console.log("Deployer balance:", stable.balanceOf(msg.sender));
    }
}
