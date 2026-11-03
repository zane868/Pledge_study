# Sepolia 测试网部署

网络：Sepolia，chainId `11155111`。部署脚本：`pledgev2/script/DeploySepolia.s.sol`。

## 账户和权限

Owner / 部署执行账户：`0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c`。费用接收地址为同一账户。
管理员采用 2/3 多签：

- `0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c`
- `0xE7018da7BD70fC8696524257550585f4166E2332`
- `0x6fA81f122444Fe58dcF6388BA2dDE91587565416`

上述三地址是多签审批者。Owner 已通过多签申请获得池、预言机和凭证权限管理的调用授权。
其他执行账户可按相同申请及审批流程获得授权；管理员身份不绕过 validCall。

## 演示池参数

- 池索引：0；原生 ETH 抵押，18 位测试稳定币借款。
- 抵押率：200%；年利率：20%；期限：7 天；两端平台费用初始为零。
- 募集截止：2026-11-04 15:43:00 UTC+8。
- 到期时间：2026-11-11 15:43:00 UTC+8。
- 最低出借限制：严格大于 1 测试稳定币。
- 初始 DEX 流动性：0.1 测试 ETH + 200 测试稳定币。
- 测试手动价格：稳定币 1、ETH 2000，不是接入实时外部价格源。

## 复现与校验

使用已配置的 `SEPOLIA_RPC_URL`、`SEPOLIA_PRIVATE_KEY`、`SEPOLIA_PRIVATE_KEY_2` 和
`ETHERSCAN_API_KEY` 环境变量，在 `pledgev2/` 执行：

```sh
make deploy-sepolia
make check-sepolia
```

部署使用 `sepolia` profile 的 Amsterdam 运行时；Solidity 0.8.24 的编译目标仍为 Paris。
Sepolia 已于 2026-10-06 激活 Glamsterdam 的新 Gas 规则，不能使用旧模型的 Gas 估算广播。
依据：[以太坊基金会公告](https://blog.ethereum.org/en/2026/09/17/glamsterdam-testnet-announcement)。

脚本模拟生成的 `deployments/sepolia.json` 不是部署完成凭证；最终地址及交易以链上校验后生成的
`pledgev2/networks/sepolia.json` 为准。

实际部署已完成：30 笔交易均成功；10 份链上运行时代码与本地编译产物匹配，
多签成员、2/3 门槛、Owner 调用授权、池参数、凭证 minter 和 DEX 储备均已核对。
Amsterdam 运行时下的 20 个测试全部通过，算术模糊测试执行 256 轮。

首次部署多签的交易因旧运行时低估 Gas 而回滚，没有留下合约；失败交易为
`0x75a567b8896b71499e195680a9f252cd5b4443ed6335962fc676f50530ce50d0`。
后续部署已切换 Amsterdam 运行时，并使用新 nonce 重新部署。

## 实际地址

| 合约 | Sepolia 地址 |
|---|---|
| PledgePool | [0x91A621F4D3f3e72b8bbD53D7387a914b9Ef29E82](https://sepolia.etherscan.io/address/0x91A621F4D3f3e72b8bbD53D7387a914b9Ef29E82) |
| 多签 | [0x7B64C65df68B5561c37c852c0cd37F4E544Ad4B7](https://sepolia.etherscan.io/address/0x7B64C65df68B5561c37c852c0cd37F4E544Ad4B7) |
| 测试预言机 | [0x72F76193bAB25b1C6d8380D09965172e6d676168](https://sepolia.etherscan.io/address/0x72F76193bAB25b1C6d8380D09965172e6d676168) |
| 测试稳定币 | [0xFfeC45EDC89BdA67A2A80AA320838a2843439ACe](https://sepolia.etherscan.io/address/0xFfeC45EDC89BdA67A2A80AA320838a2843439ACe) |
| 测试 WETH | [0xd2F3C0211E0535Ee0112D3aCB758C7429D488Af9](https://sepolia.etherscan.io/address/0xd2F3C0211E0535Ee0112D3aCB758C7429D488Af9) |
| DEX Factory | [0x0B0382c73431915EC7677b4576148a840Ec05750](https://sepolia.etherscan.io/address/0x0B0382c73431915EC7677b4576148a840Ec05750) |
| DEX Router | [0xB6885c1c99b2ce34fba4F543092840980DEC8124](https://sepolia.etherscan.io/address/0xB6885c1c99b2ce34fba4F543092840980DEC8124) |
| SP | [0x2337c550341aa23D2e359d877C7f294147E6f647](https://sepolia.etherscan.io/address/0x2337c550341aa23D2e359d877C7f294147E6f647) |
| JP | [0x72632984084dbA6B983Fd3B76E893a81CF8491b1](https://sepolia.etherscan.io/address/0x72632984084dbA6B983Fd3B76E893a81CF8491b1) |
| DEX 交易对 | [0x096c4930869da28a006a9409a26faf38b9995431](https://sepolia.etherscan.io/address/0x096c4930869da28a006a9409a26faf38b9995431) |

部署区块范围：`12053880`–`12053910`。
完整交易哈希保存在 `pledgev2/networks/sepolia.json`。

## 源码验证说明

Router 的 via-IR 编译结果受整次编译输入影响。Foundry 默认提交单文件输入时，
结果与部署字节码不同；本次用部署时 49 个源文件的完整 Standard JSON 编译输入重新核对，
创建字节码精确匹配后补交验证。输入保存于 `pledgev2/deployments/router-standard-input.json`，
其 SHA-256 保存于公开部署记录。

最终已通过 Etherscan API 逐一确认：主池、多签、预言机、测试稳定币、WETH、Factory、
Router、SP、JP 和交易对，共 10 份源码全部验证成功。

若后续执行部署命令时只发生源码验证失败，应先核对链上回执，使用同次部署记录补验，
避免误将验证失败当作部署失败而重复广播。
