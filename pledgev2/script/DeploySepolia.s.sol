// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {LocalSystem} from "./LocalSystem.sol";

contract DeploySepolia is Script, LocalSystem {
    address internal constant EXPECTED_OWNER = 0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c;
    address internal constant ADMIN_B = 0xE7018da7BD70fC8696524257550585f4166E2332;
    address internal constant ADMIN_C = 0x6fA81f122444Fe58dcF6388BA2dDE91587565416;
    mapping(address => uint256) internal keys;

    function _start(address sender) internal override {
        vm.startBroadcast(keys[sender]);
    }

    function _stop() internal override {
        vm.stopBroadcast();
    }

    function _stableUnit() internal pure override returns (uint256) { return 1e6; }
    function _stableArtifact() internal pure override returns (string memory) { return "out/TestUSDT.sol/TestUSDT.json"; }

    // 固定复用已部署的 TestUSDT，不再重复铸币。
    function _stableAddress() internal pure override returns (address) {
        return 0x0Cc069Daa724C367e0210B60994A197B05a80aBf;
    }

    function _liquidityEth() internal pure override returns (uint256) {
        return 0.1 ether;
    }

    function _poolDuration() internal pure override returns (uint256) {
        return 7 days;
    }

    function _poolDelay() internal pure override returns (uint256) {
        return 1 days;
    }

    function _signatureOwners() internal view override returns (address[] memory owners) {
        owners = new address[](3);
        owners[0] = ownerA;
        owners[1] = ownerB;
        owners[2] = ADMIN_C;
    }

    function _key(string memory variableName) internal view returns (uint256) {
        string memory value = vm.envString(variableName);
        if (bytes(value).length == 64) value = string.concat("0x", value);
        return vm.parseUint(value);
    }

    function run() external {
        require(block.chainid == 11155111, "Sepolia only");
        uint256 deployerKey = _key("SEPOLIA_PRIVATE_KEY");
        uint256 cosignerKey = _key("SEPOLIA_PRIVATE_KEY_2");
        operator = vm.addr(deployerKey);
        ownerA = operator;
        ownerB = vm.addr(cosignerKey);
        require(ownerA == EXPECTED_OWNER && ownerB == ADMIN_B, "Configured keys do not match owners");
        keys[ownerA] = deployerKey;
        keys[ownerB] = cosignerKey;
        lender = operator;
        borrower = ownerB;
        recipient = ownerB;
        fee = operator;

        _deploySystem();
        _start(operator);
        pool.setMinAmount(1e6); // Six-decimal test USDT; deposit must be strictly greater.
        _stop();

        string memory name = "sepolia";
        vm.serializeString(name, "compiler", "0.8.24");
        vm.serializeString(name, "environment", "testnet");
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
        vm.serializeAddress(name, "admins", _signatureOwners());
        vm.serializeUint(name, "threshold", 2);
        vm.serializeAddress(name, "feeAddress", fee);
        vm.serializeUint(name, "stableDecimals", 6);
        vm.serializeUint(name, "initialLiquidityEthWei", _liquidityEth());
        vm.serializeUint(name, "poolId", pid);
        vm.serializeUint(name, "settleTime", settleTime);
        string memory manifest = vm.serializeUint(name, "endTime", endTime);
        vm.createDir("deployments", true);
        vm.writeJson(manifest, "deployments/sepolia.json");
        console2.log("PledgePool", address(pool));
        console2.log("Multisig", address(multisig));
        console2.log("Test stablecoin", address(stable));
    }
}
