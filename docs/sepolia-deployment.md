# Sepolia ETH / USDT 测试部署

2026-11-03 更新为 USDT 六位精度版本，chainId 为 `11155111`。
Owner 和部署执行账户：`0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c`。
管理员为已确认的三地址，门槛 2/3，见部署清单中的 admins 字段。

| 合约 | Sepolia 地址 |
|---|---|
| PledgePool | [0x1D462BBE7451F555c6742E09972a61AFFfEaD1e9](https://sepolia.etherscan.io/address/0x1D462BBE7451F555c6742E09972a61AFFfEaD1e9) |
| 2/3 多签 | [0xAdC32122226C387CDF0b602bD76Ccc32041eFC8F](https://sepolia.etherscan.io/address/0xAdC32122226C387CDF0b602bD76Ccc32041eFC8F) |
| 测试预言机 | [0xF691AA22Da250f32f55afD139491a0F44C21c259](https://sepolia.etherscan.io/address/0xF691AA22Da250f32f55afD139491a0F44C21c259) |
| Test USDT（6 位） | [0x9DB37141f5708ADA96AeFE69ce84623880Dfe184](https://sepolia.etherscan.io/address/0x9DB37141f5708ADA96AeFE69ce84623880Dfe184) |
| 测试 WETH | [0x841E32459367af4D91C5810dea7979Bf1C42e578](https://sepolia.etherscan.io/address/0x841E32459367af4D91C5810dea7979Bf1C42e578) |
| DEX Factory | [0xfE423Fdf5fc9C6df85116F42C036E7e347C5701F](https://sepolia.etherscan.io/address/0xfE423Fdf5fc9C6df85116F42C036E7e347C5701F) |
| DEX Router | [0xE64974473A2Fc860De6D594d37Ad92846b45CB99](https://sepolia.etherscan.io/address/0xE64974473A2Fc860De6D594d37Ad92846b45CB99) |
| SP（18 位） | [0x8c2094896f8B2bA54d663CE7Acd1294447e0F164](https://sepolia.etherscan.io/address/0x8c2094896f8B2bA54d663CE7Acd1294447e0F164) |
| JP（18 位） | [0x210bbD576f78da5D1645fcDc8f7325621d89A90f](https://sepolia.etherscan.io/address/0x210bbD576f78da5D1645fcDc8f7325621d89A90f) |
| DEX 交易对 | [0x1bfd0b85c162aa9821d644806398a3d6c3dc0f35](https://sepolia.etherscan.io/address/0x1bfd0b85c162aa9821d644806398a3d6c3dc0f35) |

31 笔部署及初始化交易成功，10 份源码全部经浏览器验证。
链上代码、合约关联、权限、价格、储备、USDT 6 位、凭证 18 位及 lendScale=10^12 均已核对。
部署区块：`12055166`–`12055197`。

- ETH 抵押、测试 USDT 借款；Test USDT 不是发行方官方资产。
- 初始流动性：0.1 测试 ETH 与 200 测试 USDT。
- 抵押率 200%，年利率 20%，募集一天，匹配后期限七天；平台费用初始为零。
- 最低出借限制为严格大于 1 USDT，即原始金额必须大于 1000000。
- 预言机是手动测试价格：ETH=2000、USDT=1，均以 10^8 表示。

完整地址、时间戳、回执和验证状态保存于 `pledgev2/networks/sepolia.json`。
本项目按全新 ETH/USDT 系统处理，不提供旧资产兼容或迁移流程。

## 命令

```sh
cd pledgev2
FOUNDRY_PROFILE=sepolia forge test -vv
make deploy-sepolia
make check-sepolia
```

部署需要环境中的 RPC 和两位签署者私钥；后端仅需 RPC，不需要私钥。
运行时使用 Amsterdam Gas 规则，编译目标保持 Paris。
只在创建新部署时执行 deploy，源码验证失败不应触发重复部署。
需要完整编译输入时，部署命令加 `--build-info`，再使用 `tools/verify-exact-build.py` 补验。
本次完整编译信息保存在本地忽略的 `out/build-info/`。

精度与变更规则见 [USDT6 迁移记录](usdt6-migration.md)。
