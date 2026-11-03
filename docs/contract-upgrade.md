# Solidity 0.8.24 升级记录

日期：2026-11-03。源版本：`ProjectBreakdown-Pledge` 提交
`78b418d2f5e7acde8881e142d6369fc01186ff9e` 的 `pledgev2/contracts`。

## 回退及工具链

- 撤回旧版本编译器兼容方案，先恢复源合约，再在此基础上升级。
- 移除 Hardhat 配置、JavaScript 部署和测试、npm 包及锁文件、solc-js 启动器和旧构建产物。
- Foundry 使用单一 solc 0.8.24；直接执行 `forge build`、`forge test`、`forge script`。
- 优化器开启，runs 为 1；启用 `via_ir`，EVM 目标为 Paris。
- OpenZeppelin 从 3.4.2 升为 4.9.6；forge-std 固定为 1.9.7。
  依赖源码随项目保存，不需要 npm 安装。
- PRD 的用户已有修改保留，源项目目录未修改。

## 合约改动

| 范围 | 具体变化 | 目的与行为影响 |
|---|---|---|
| 全部 24 个合约、库和接口 | pragma 统一为 `0.8.24` | 消除 0.4、0.5、0.6 多编译器依赖 |
| PledgePool、MockPledgePool | ReentrancyGuard 导入改为 `security/ReentrancyGuard.sol` | 适配 OpenZeppelin 4.9.6 |
| AddressPrivileges | EnumerableSet 导入改为 `utils/structs/EnumerableSet.sol` | 适配 OpenZeppelin 新目录 |
| 各构造函数 | 删除 `public`、`internal` 可见性声明 | 适配新版构造函数语法，保留参数和基类初始化 |
| 池、预言机、多签存储 | 地址与整数转换显式经过 `uint160` | 保留地址的 160 位表示；多签存储位置不变 |
| 用户资金支付及测试代币 | 显式使用 `payable(msg.sender)` | 适配 ETH 接收地址类型检查 |
| 池和 DEX 的授权 | `uint(-1)` 等改为 `type(uint256).max` 或对应类型最大值 | 保留无限授权数值 |
| 池兑换 | `now` 改为 `block.timestamp` | 保留执行时 `timestamp + 30` 的原期限规则 |
| 中文错误信息 | 改为 Unicode 字符串字面量 | 保留原有错误文案，满足 0.8 字符串语法 |
| WETH9 测试合约 | 匿名 payable 函数改为 `receive()`；补 `emit`；`this.balance` 改为 `address(this).balance` | 保留包装、解包、转账和授权接口 |
| DEX factory、LP 和 router | 整理公共 ERC20 接口继承，删除实现中重复事件，补 getter 的 `override` | 解决新版继承与事件重复检查；事件和外部函数签名继续可用 |
| DEX 接口只读性 | 工厂、WETH、initCodeHash 等 getter 使用 `view`，pairFor 对应改为 `view` | 正确描述读取状态的函数；CREATE2 继续读取本次编译的 initCodeHash |
| DEX 地址计算与汇编 | CREATE2 地址转换显式经过 `uint160`；`chainid` 改为 `chainid()` | 适配新版类型与汇编语法 |
| 自带 SafeMath 库 | 在带有原始显式校验的算术函数内局部使用 `unchecked` | 保留原有溢出保护及错误字符串；不会返回未检查的溢出结果 |
| DEX TWAP | uint32 时间差和 uint256 累计价格加法局部使用 `unchecked` | 保留 Uniswap V2 的时间戳及累计价格模运算，避免升级后跨界回滚 |

其他原始算术现在遵循 Solidity 0.8 的默认溢出检查，异常路径可能使用 Panic 错误。
本次保留募集、匹配、比例分配、固定分母、期限计息和费用公式；没有趁版本迁移
调整清算、损失分配或治理规则。OpenZeppelin ERC20 的失败信息及内部实现随依赖升级变化。

## 验证

在 `pledgev2/` 执行：

```sh
forge build --sizes
forge test -vv
```

编译通过。20 个测试通过，0 失败；其中算术模糊测试运行 256 轮。
覆盖真实 2/2 多签与正式池、权限拒绝、原生 ETH 入金、多人匹配、退款、借款和凭证领取、
募集失败退出、DEX 兑换与费用、短期计息、凭证转让和部分赎回、迟领、暂停、清算和失败回滚。
升级专项测试确认 SafeMath 错误字符串，以及 DEX 时间戳和价格累计值跨最大值后仍正确回绕。
另已通过 Anvil 本地广播部署和 `make abi` 导出验证。

原始和升级后文件 SHA-256 同时记录在 `pledgev2/reference/source-manifest.json`。
原项目测试保留在 `reference/legacy-tests/`，不参与 Forge 验收。

Foundry 仍会提示源实现中的事件顺序、转账、精度与类型转换等 lint 项；编译通过不等于
这些现有设计问题已经解决。PRD 中的账户级授权复用、清算阈值、滑点保护、池级预算、
价格有效性和不足额退出仍见 [合约实施记录](contract-baseline.md)。

后续已实施 USDT 六位精度适配，现行规则及新增测试见 [USDT6 迁移](usdt6-migration.md)。
