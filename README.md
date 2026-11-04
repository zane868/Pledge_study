# Pledge study

产品需求见 [Pledge_PRD.md](Pledge_PRD.md)。当前先建设本地核心合约闭环。

合约从 `ProjectBreakdown-Pledge/pledgev2/contracts` 迁入后统一升级为 Solidity 0.8.24，
使用 Foundry 构建、测试和本地部署。前端、Go 服务和自动执行服务属于后续阶段。

```sh
cd pledgev2
forge build
forge test -vv
```

运行说明见 [合约工程](pledgev2/README.md)，具体改动见
[Solidity 升级记录](docs/contract-upgrade.md)，PRD 实现边界见
[合约实施记录](docs/contract-baseline.md)。

测试网部署及权限配置见 [Sepolia 部署说明](docs/sepolia-deployment.md)。

Go 后端已迁入 [pledge-backend](pledge-backend/README.md)，迁移范围及 PRD 差异见
[后端功能对照](docs/backend-migration.md)，服务器部署见
[后端服务器部署](docs/backend-server-deployment.md)。

用户端前端在 [pledge-fe](pledge-fe/)，与前端共用的链上读数和数值计算在
[shared](shared/)。前端已部署到 <http://43.160.212.151/app/>，构建与发布流程见
[前端服务器部署](docs/frontend-server-deployment.md)。

```sh
cd pledge-fe
npm ci
npm run build     # 资源前缀固定为 /app/
```

`shared/src/numbers.ts` 复刻了合约的取整顺序，是「预计可领 / 预计退款 / 预计赎回」
的唯一实现来源。它的测试 `shared/src/numbers.test.ts` 与合约侧基准测试
`pledgev2/test/PledgeRounding.t.sol` 使用同一组向量，两侧必须同步修改。
