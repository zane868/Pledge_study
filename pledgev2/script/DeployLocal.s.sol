// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {LocalSystem} from "./LocalSystem.sol";

contract DeployLocal is Script, LocalSystem {
    mapping(address => uint256) internal keys;

    function _start(address sender) internal override {
        vm.startBroadcast(keys[sender]);
    }

    function _stop() internal override {
        vm.stopBroadcast();
    }

    function _account(uint32 index) internal returns (address account) {
        // Public Anvil test mnemonic, used exclusively on a local chain.
        uint256 key = vm.deriveKey("test test test test test test test test test test test junk", index);
        account = vm.addr(key);
        keys[account] = key;
    }

    function run() external {
        require(block.chainid == 31337, "Local Anvil chain only");
        operator = _account(0);
        ownerA = _account(1);
        ownerB = _account(2);
        lender = _account(3);
        borrower = _account(4);
        recipient = _account(5);
        fee = _account(6);
        _deploySystem();

        string memory name = "local";
        vm.serializeUint(name, "chainId", block.chainid);
        vm.serializeAddress(name, "multisig", address(multisig));
        vm.serializeAddress(name, "oracle", address(oracle));
        vm.serializeAddress(name, "stable", address(stable));
        vm.serializeAddress(name, "weth", weth);
        vm.serializeAddress(name, "factory", factory);
        vm.serializeAddress(name, "router", address(router));
        vm.serializeAddress(name, "pool", address(pool));
        vm.serializeAddress(name, "sp", address(sp));
        vm.serializeAddress(name, "jp", address(jp));
        vm.serializeAddress(name, "collateral", address(0));
        vm.serializeAddress(name, "operator", operator);
        vm.serializeAddress(name, "ownerA", ownerA);
        vm.serializeAddress(name, "ownerB", ownerB);
        vm.serializeAddress(name, "feeAddress", fee);
        vm.serializeUint(name, "stableDecimals", 18);
        vm.serializeUint(name, "poolId", pid);
        vm.serializeUint(name, "settleTime", settleTime);
        string memory manifest = vm.serializeUint(name, "endTime", endTime);
        vm.createDir("deployments", true);
        vm.writeJson(manifest, "deployments/local.json");
        console2.log("PledgePool", address(pool));
        console2.log("SP", address(sp));
        console2.log("JP", address(jp));
    }
}
