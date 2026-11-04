# @pledge/shared

`pledge-fe` 与后续 `admin-fe` 共用的链上读数和数值计算。

## 内容

| 文件 | 作用 |
|---|---|
| `src/numbers.ts` | 复刻合约的匹配、凭证分母、退款、赎回、计息等算法 |
| `src/poolState.ts` | 池的五态定义与页面派生状态（PRD 5.3.1 / 5.3.4） |
| `src/config.ts` | 部署清单的类型定义与地址辅助函数 |
| `src/api.ts` | 后端 API 客户端 |
| `deployment.json` | 当前 Sepolia 部署的地址清单，由 `pledgev2/networks/sepolia.json` 对齐 |
| `abi/` | 从 `pledgev2/abi/` 同步的 ABI，不入库，用 `npm run sync-abi`（在 pledge-fe 下）刷新 |

## 数值层为什么重要

链上金额全部是整数，取整方向和乘除顺序直接决定实际到账数量。前端如果用浮点、
或者把链上分两步的除法合并成一步，用户看到的「预计可领」就会和实际不一致。

`src/numbers.ts` 的每个函数都对应合约里的具体分支，注释里标了出处。
`src/numbers.test.ts` 用与合约侧 `pledgev2/test/PledgeRounding.t.sol` 同一组
向量做断言 —— 两边必须同步修改，否则测试会失败。

## 测试

```sh
npm install
npm test
```
