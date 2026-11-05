# 已知问题与待办清单

更新时间：2026-11-05。本文件汇总当前会话确认的待处理问题，按模块分组，供后续接手继续处理。

---

## 一、运营后台「创建池子」向导

代码：`admin-fe/src/pages/CreatePoolPage.tsx`

| # | 问题 | 现状 | 建议 |
|---|---|---|---|
| 1 | **进度不持久化** | 向导状态（step / params / sp / jp 地址）全在 `useState` 内存，关闭或刷新即丢。已实际发生过一次重复部署 SP（链上多出一个废弃的 SP 合约，白花一笔 gas）。目前只在页面顶部加了风险提示卡片。 | 增加 `localStorage` 持久化；或提供「手动填入已部署地址」恢复入口。 |
| 2 | **已部署凭证无法恢复** | 页面中断后没有地方输入已部署的 SP/JP 地址，只能重新部署（产生额外合约创建 gas）。 | 在第二步增加「我已部署过，填入地址」输入框。 |
| 3 | **多签签名需手动切钱包** | `useMultisigApproval` 已固定申请人为 operator（修复了切换管理员钱包后申请「消失」的 bug），并新增了 `/admin/sign?to=凭证&label=SP` 独立签名页 + 复制链接。但签名仍是链上交易，必须由管理员本人钱包发起，无法代签。 | 体验上已尽量优化；如需进一步简化可考虑批量签名或离线签名方案。 |
| 4 | **创建成功后重复建池风险** | 曾有「交易成功后页面无反馈 → 用户重复点击」导致创建出两个参数完全相同的池（pid 1、pid 2）。已增加 `receipt.status === 'success'` 成功态页面 + 按钮禁用。遗留的重复池无法删除（合约无删池入口）。 | 链上遗留池只能 `forceSettle` 空池转 UNDONE；根本解决需要合约增加池撤销能力或前端幂等校验。 |
| 5 | **缺少创建后回读核对** | PRD 6.4.1 / 6.4.2 要求交易确认后回读 `poolLength`、新池参数与状态并逐项核对，当前前端无此环节。 | 创建成功后展示链上回读值与提交值的对比。 |

---

## 二、合约层（P0 资金与权限）

### 2.1 多签授权不可消耗，一次授权 = 无限次使用

`multiSignatureClient.validCall` 只校验 `keccak256(调用者, 合约地址)` 是否被多签批准，**不区分函数、参数、调用次数、有效期**。

一旦某个申请人（如部署时的 operator）对池合约的授权达标，他就可以无限次调用该合约**所有** `validCall` 函数：

- `createPoolInfo`（可创建 N 个池子）
- `settle` / `forceSettle`
- `finish` / `forceFinish`
- `liquidate` / `forceLiquidate`
- `setFee` / `setFeeAddress` / `setMinAmount` / `setSwapRouterAddress`
- `setPoolSettleTime` / `setPoolEndTime` / `setPoolState`（本次新增）

直到管理员调用 `revokeSignApplication` 撤销签名。

**风险**：一次授权即等于该合约的永久全权限。正确做法应类似 Gnosis Safe：每笔操作（含函数、参数、nonce）独立提案，执行一次即消耗，并可设过期时间。

依据：`pledgev2/contracts/multiSignature/multiSignatureClient.sol`、`docs/handoff/2026-11-05-handoff.md` §5.1。

**影响面**：创建池子、结算、清算、费率变更等所有管理操作都受此设计影响。

### 2.2 清算/结算卖出的滑点保护缺失

- `_swap` 里 `minOut` 固定为 `0`，无最小产出保护
- `deadline` 在执行时才生成 `block.timestamp + 30`，无报价有效期
- 先 `getAmountsIn` 反推输入、再按输入精确卖出，实际到账可能少于应还本息

### 2.3 其他已记录合约问题

见 `docs/HANDOFF.md` §10 与 `docs/handoff/2026-11-05-handoff.md`，主要包括：

- 清算入口未强制风险阈值
- 提前清算仍按完整约定期限计息
- 全局费率/路由变更影响存续池，未冻结费用快照
- ETH 抵押侧无最低门槛（`minAmount` 只作用于出借侧）

---

## 三、用户端「出借」需两次交互

代码：`pledge-fe/src/pages/PoolDetailPage.tsx`

当前出借流程是两步独立交易：

1. `approve` — 授权 USDT 给池合约
2. `depositLend` — 实际入金

钱包需要确认两次。

`TestUSDT`（`0x0Cc069...`）已继承 `ERC20Permit`，支持 EIP-2612，但**合约和前端均未接入 permit**。

| 方案 | 改动 | 效果 |
|---|---|---|
| 短期：前端 max 授权 | 仅前端（approve 改为 `max uint256`） | 首次两步，之后每次一步 |
| 长期：permit 单笔入金 | 合约加 `depositLendWithPermit` + 前端 `signTypedData` | 每次出借都一次交互 |

---

## 四、平台收益未在页面体现

目前 `lendFee` / `borrowFee` 均为 0，平台无收益；且前端/后台**都没有展示**：

- 当前费率（用户出借时不知情）
- 实际收取的收益金额
- `feeAddress`（收益接收地址）

PRD 6.3.2 要求「页面需分别显示借款利息、平台费用、兑换成本与 Gas」。若要启用收费，需要同步补齐展示与管理入口（`setFee`）。

---

## 五、运营后台登录方式（已调整）

2026-11-05 调整：页面访问改为**账号密码登录**（后端 `/user/login`，JWT 有效期 12 小时），涉及链上写操作时才需要连接管理员钱包。

已知遗留：

- 目前只有创建池页接入了 `WalletGate` 检测；池管理页的结算/签名按钮、`SignPage` 尚未统一接入钱包就绪检测，未连钱包时点击可能直接报错。

---

## 附：当前部署地址（易变，以链上为准）

| 合约 | 地址 |
|---|---|
| PledgePool | `0x9A003A3fF42927c56E17593BFe4608Cab60fECB7` |
| TestUSDT（固定复用） | `0x0Cc069Daa724C367e0210B60994A197B05a80aBf` |
| multisig | `0x405BA09EB72ca910090a993a134842452fB258C4` |
| Oracle | `0xD229e4C7F4EbEfAF7b382D771Db0e2ADC7F067b5` |
| SP / JP | `0xBbC9b8a5dF325a7C95266fad8ed88A24fB6a666a` / `0x0621B4B47ADc96d25F29967f0275218e26fed92E` |

> 地址随时可能因重新部署变化，以 `pledgev2/networks/sepolia.json` 与链上数据为准。
