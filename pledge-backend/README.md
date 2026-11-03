# Pledge Go 后端

从 `ProjectBreakdown-Pledge/pledge-backend` 迁入，保留 API、模型、工具与调度目录。
本次将活动链改为本项目已部署的 Sepolia，并修正核心同步和查询问题。
原始配置凭证、数据库种子及宿主机 service 文件未迁入。

## 运行

需要 Go 1.25+；当前使用 Go 1.27.0 验证。运行服务需要 MySQL、Redis，链上同步另需
`SEPOLIA_RPC_URL`。不需要部署账户私钥。

```sh
make build
make test
make chaincheck
```

`chaincheck` 不依赖数据库，只读核对 RPC 网络、确认区块、池、费用、资产与预言机价格。
活动部署来自 `config/sepolia.json`，由仓库的 `pledgev2/networks/sepolia.json` 生成。

准备配置：

```sh
cp .env.example .env
```

填入数据库密码、RPC、后台密码（至少 12 位）和 JWT secret（至少 32 位）。
`.env` 不进入 Git。Go 进程不会自动加载 `.env`；直接运行时先导出其中的环境变量。

已有本地 MySQL/Redis 时分别启动：

```sh
make api
# 另一终端，使用相同环境配置
make sync
```

也提供容器配置：

```sh
docker compose up --build
```

API 默认监听 8080，仅容器端口映射到本机回环地址。数据卷保留 MySQL 数据。
同步每两分钟执行一次，读取最新高度减 12 个确认块的完整快照。SQL 提交后才更新 Redis，
启动不会清空 Redis，也不会自动发送邮件、订阅 PLGR 行情或签署链上交易。

## 接口

版本前缀为 `/api/v1`，保留原路由命名和 `code/message/data` 兼容格式。

| 接口 | 行为 |
|---|---|
| `GET /healthz` | MySQL、Redis 就绪检查 |
| `GET /api/v1/deployment` | 部署标识、合约地址、Owner、三位管理员和门槛 |
| `GET /api/v1/sync/status` | 已发布区块、完整性、最后成功及错误和过期标记 |
| `GET /api/v1/poolBaseInfo?chainId=11155111` | 当前部署的池基础信息 |
| `GET /api/v1/poolDataInfo?chainId=11155111` | 按池隔离的匹配和终态分配数据 |
| `GET /api/v1/token?chainId=11155111` | 已同步资产元数据 |
| `POST /api/v1/user/login` | JSON `name/password` 登录 |
| `POST /api/v1/user/logout` | 注销当前后台会话 |
| `POST /api/v1/pool/search` | 鉴权后按 `chainID`、资产、状态分页查询 |
| `POST /api/v1/pool/debtTokenList` | 保留来源项目的代币元数据查询接口 |
| `POST /api/v1/pool/setMultiSign`、`getMultiSign` | 读写后台配置记录，不替代链上多签审批 |
| `GET /api/v1/price` | 原 PLGR 价格功能未启用，返回 503 |

后台请求通过 `authCode` header 传入登录所得 token。后台登录身份与钱包、链上多签权限分离。
无成功同步、上轮同步失败或快照超过五分钟时，池查询返回 503，避免呈现虚假的空市场。
历史兼容接口仍可能使用 HTTP 200 携带业务错误，客户端同时检查 `code`。

金额与比例保持十进制原始整数字符串。利率、费用使用 `10^8` 基数；当前 Sepolia 手动预言机
价格以 `10^8` 表示。ETH 为零地址，SP/JP 不假定具有预言机价格。
链上索引 `poolIndex=0` 对应 `displayPoolId=1`；返回值带池合约及同步区块。

## 绑定与依赖

先在 `../pledgev2` 执行 `forge build`，安装与模块匹配的 abigen 1.17.8，然后执行：

```sh
make bindings
```

绑定包含当前 PledgePool、Oracle、DebtToken 和 MultiSignature ABI。
活动网络固定为 Sepolia，借款资产为六位测试 USDT。部分来源类型名仍带旧命名，不表示支持其他网络或资产。
依赖版本由 `go.mod/go.sum` 锁定。

测试使用 SQLite 和进程内 Redis 验证数据及 HTTP 行为；实际部署继续使用 MySQL 和 Redis。
本地未安装 Docker/MySQL/Redis。已在 Ubuntu 服务器使用实际 MySQL/Redis 完成独立服务部署与联调，见
[服务器部署说明](../docs/backend-server-deployment.md)；Docker 构建仍未实测。
PRD 的完整交易历史索引、重组回滚、仓位发现、自动执行和审计尚未全部完成，
见 [后端迁移与功能对照](../docs/backend-migration.md)。
