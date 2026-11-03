# Pledge 合约工程

使用 Solidity **0.8.24**、OpenZeppelin **4.9.6** 和 Foundry。
依赖源码已放在 `lib/`，版本记录见 [dependencies.json](reference/dependencies.json)。
构建使用单一原生 solc，不需要 npm、Hardhat 或旧编译器启动器。

## 构建与测试

在已安装 Foundry 的环境执行：

```sh
forge build
forge test -vv
```

Foundry 首次运行会获取 solc 0.8.24。配置启用优化器、`via_ir`，EVM 目标为 Paris。
原生合约逻辑和本地 DEX、测试资产均升级到同一编译器版本。

测试覆盖正式多签、权限、ETH 入金、比例匹配、退款、领取、转让、部分赎回、
到期费用、短期计息、募集失败、暂停、清算和失败回滚，另外验证 DEX 时间戳与
累计价格的回绕语义，以及 SafeMath 的原有错误信息和模糊测试。
`reference/legacy-tests/` 仅保留来源测试作参考，不参与 Forge 测试。

## 本地部署

两个终端分别执行：

```sh
anvil
```

```sh
forge script script/DeployLocal.s.sol:DeployLocal \
  --rpc-url http://127.0.0.1:8545 --broadcast
```

脚本限制 chainId 为 31337，并使用 Anvil 默认公开测试助记词派生账户。
它部署正式池、预言机、2/2 多签和独立 SP/JP，以及测试稳定币、WETH 和 DEX 流动性；
完成多签签署与池 minter 授权。仅用于本地测试环境。

测试池以 ETH 抵押，借出 18 位测试稳定币；示例抵押率 200%、年利率 20%、期限 365 天。
地址和池参数写入 `deployments/local.json`，广播回执写入 `broadcast/`。
本地链重置后重新部署。

导出当前 ABI：

```sh
make abi
```

ABI 从 `out/` 编译产物生成到 `abi/`。`Makefile` 另提供 build、test、fmt、node、deploy-local 入口。

## Sepolia 部署

环境变量提供 `SEPOLIA_RPC_URL`、`SEPOLIA_PRIVATE_KEY`、`SEPOLIA_PRIVATE_KEY_2`；
源码验证使用 `ETHERSCAN_API_KEY`。私钥只从环境读取，不写入项目文件。

```sh
make deploy-sepolia
make check-sepolia
FOUNDRY_PROFILE=sepolia forge test -vv
```

`sepolia` profile 使用 Amsterdam 运行时计价，编译目标继续保持 Paris。
管理员为指定的三个地址，签署门槛为 2/3，部署和管理执行账户为指定 Owner。
脚本注入 0.1 测试 ETH 和 200 测试 USDT流动性，募集一天，匹配后期限七天；
最低出借限制为严格大于 1 测试稳定币。预言机为测试手动价格，不是正式价格源。

`deployments/sepolia.json` 是脚本运行清单，模拟执行也会生成它。
实际广播完成后，检查脚本核对回执、链上字节码和权限，再生成可保存的
`networks/sepolia.json`。链上部署记录及地址见 [测试网部署说明](../docs/sepolia-deployment.md)。

## 改动与边界

升级清单见 [contract-upgrade.md](../docs/contract-upgrade.md)。
本次保留匹配、计息、费用和清算的原业务规则，PRD 中的既有差异见
[contract-baseline.md](../docs/contract-baseline.md)。
`reference/source-manifest.json` 同时记录源版本文件哈希与升级后文件哈希。

当前测试网已完成 ETH/USDT 六位精度适配，SP/JP 保持 18 位，详见 [USDT6 迁移](../docs/usdt6-migration.md)。
