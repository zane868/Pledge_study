// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {PledgePool} from "../contracts/pledge/PledgePool.sol";

/**
 * @notice 重新部署 PledgePool（加入 forceSettle / forceFinish 的版本）。
 *
 * @dev PledgePool 不是可升级合约（无 proxy），因此要上线新接口必须重新部署。
 *      本脚本只做「部署新池 + 复用既有外部依赖」，不迁移链上已有池子的资金：
 *
 *      - 新的 PledgePool 复用同一 oracle / swapRouter / multiSignature / feeAddress；
 *      - 旧的 PledgePool 合约保持不变，旧池（pid=0）及其 500 USDT 仍在旧合约上；
 *      - 新合约从 pid=0 开始重新创建池子。
 *
 *      重要：旧池的出借人需要在旧合约上先 refundLend 取回资金，
 *      否则那笔 USDT 会留在旧合约里无法在新系统使用。
 *      旧池处于 MATCH 且 borrowSupply=0，执行 settle 会进入 UNDONE，
 *      之后可用 emergencyLendWithdrawal 取回本金（需池状态为 UNDONE）。
 */
contract RedeployPledgePool is Script {
    // 已部署的外部依赖（Sepolia）
    address internal constant ORACLE = 0xF691AA22Da250f32f55afD139491a0F44C21c259;
    address internal constant SWAP_ROUTER = 0xE64974473A2Fc860De6D594d37Ad92846b45CB99;
    address internal constant MULTISIG = 0xAdC32122226C387CDF0b602bD76Ccc32041eFC8F;
    address payable internal constant FEE_ADDRESS = payable(0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c);

    function run() external {
        require(block.chainid == 11155111, "Sepolia only");
        uint256 deployerKey = vm.envUint("SEPOLIA_PRIVATE_KEY");

        vm.startBroadcast(deployerKey);
        PledgePool pool = new PledgePool(ORACLE, SWAP_ROUTER, FEE_ADDRESS, MULTISIG);
        vm.stopBroadcast();

        console2.log("New PledgePool", address(pool));
        console2.log("oracle        ", ORACLE);
        console2.log("swapRouter    ", SWAP_ROUTER);
        console2.log("multiSignature", MULTISIG);
        console2.log("feeAddress    ", FEE_ADDRESS);
        console2.log("poolLength    ", pool.poolLength());
        console2.log("minAmount     ", pool.minAmount());
    }
}
