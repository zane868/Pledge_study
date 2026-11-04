# pledge-fe

Pledge 用户端前端。React 18 + TypeScript + Vite + wagmi/viem。

## 路由

| 路径 | 页面 |
|---|---|
| `/` | 借贷市场列表 |
| `/pool/:chainId/:poolIndex` | 池详情与入金（出借 / 抵押两个 Tab） |
| `/portfolio` | 我的仓位 |
| `/portfolio/:chainId/:poolIndex` | 仓位详情与领取、退款、赎回 |

部署在 `/app/` 子路径下，见 [前端服务器部署](../docs/frontend-server-deployment.md)。

## 开发

```sh
npm ci
npm run dev        # http://localhost:3000，/pledge 代理到线上后端
npm run typecheck
npm run build
```

## 代码边界

- **数值计算**一律走 `@pledge/shared`，不在页面里手写乘除。链上取整顺序见
  `shared/src/numbers.ts` 的注释。
- **ABI 来源**是 `pledgev2/abi/`，同步到 `shared/abi/` 后用 `npm run sync-abi`
  刷新；不要在前端手改 ABI。
- **状态文案**由 `shared/src/poolState.ts` 的 `derivePoolState` 统一生成，
  页面不各自拼装。链上状态只有五类，不合并成「已结束」。

## 与 PRD 的对应

| 页面 | PRD 依据 |
|---|---|
| 市场列表 | 11.2（状态与时间展示、池定位与数据一致性） |
| 池详情与入金 | 11.3（输入校验、稳定币授权与 ETH 支付、确认与后续操作） |
| 我的仓位 | 11.4（原始参与与凭证持仓分开核算，覆盖纯凭证持有人） |
| 交易状态 | 11.5.1（等待钱包确认 / 用户拒绝 / 结果未知 / 待确认 / 失败） |

页面里已经落实的约束：

- 入金按钮的阻塞原因按 11.5.2 显示最直接的一条
- 授权与入金是两笔交易，授权成功不显示为已参与募集
- 系统时间仅作辅助，可执行性以链上最新区块时间为准（5.3.3）
