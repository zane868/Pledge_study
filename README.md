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
