# Pledge 项目交接文档

整理日期：2026-11-03（Asia/Shanghai）。本文件以本次会话最终确认和实际部署状态为准。
服务器位置：`/opt/pledge/handoff/HANDOFF.md`。

## 1. 最终产品约定

- 项目是 ETH 抵押借款：出借人提供 USDT，借款人抵押原生 ETH，募集匹配后领取 USDT。
- USDT 为 6 位；ETH 为 18 位；SP、JP 均为 18 位。比例、年利率、抵押率、费用基数为 10^8。
- 当前是 Sepolia 测试环境，自建 TestUSDT，不是发行方官方 USDT，也不是主网部署。
- 按全新系统建设，不再做旧 BUSD 兼容、存量资金迁移、旧部署归档或版本切换流程。
- 到期/清算时兑换 ETH 准备偿付，SP 赎回 USDT，JP 赎回剩余 ETH。没有主动还款取回全部原始 ETH 的现成流程。
- 合约用 Foundry；后端是 Go + MySQL + Redis；前端尚未建设。
- 用户确认无需域名。目前不配置 HTTPS，公网仅开放只读接口，后台操作通过 SSH 隧道。

## 2. 本次完成的工作

1. 阅读并保存 PRD，完成过 Git 提交；从源项目迁入合约，再统一升级 Solidity 0.8.24。
2. 使用 Foundry 替代 Hardhat；配置优化器 runs=1、via_ir，编译目标 Paris。
3. 适配 ETH 18 / USDT 6 的匹配、风险估值、铸造、费用、利息与赎回；重建测试资产、池及凭证。
4. 完成当前 Sepolia 部署：31 笔交易成功，10 份链上运行时代码核对，10 份源码在浏览器验证。
5. 从源项目迁入 Go 后端；修复核心同步、查询和登录问题，生成当前 ABI 的 Go 绑定。
6. 检查服务器环境，部署 Go API、只读同步进程、独立数据库和 Redis，保留服务器原有应用。
7. 后端已读取和同步当前六位 USDT 部署；清除了 Pledge 自身旧空测试记录和对应旧缓存。

## 3. 源码与 Git 状态（接手前先看）

原本机工作目录：`/Users/jiandong/github_space/Pledge_study`。
参考源码：`/Users/jiandong/github_space/ProjectBreakdown-Pledge`，参考提交
`78b418d2f5e7acde8881e142d6369fc01186ff9e`。
远程仓库：`git@github.com:zane868/Pledge_study.git`，当前本地分支 main。

已完成的本地提交：
- `69071f8`：PRD 第 3–9 章，已在当时推送。
- `1fb19bb`：Solidity 0.8 / Foundry 初始合约基线，已本地提交，当时未推送。

**当前 USDT6 合约改造、后端源码及运维材料仍含未提交修改和未跟踪文件。**
接手者从 Git 获取源码：

```sh
git clone git@github.com:zane868/Pledge_study.git
# HTTPS 地址：https://github.com/zane868/Pledge_study
```

本次交接不在服务器保存源码快照。当前远程仓库尚未包含本会话全部最新实现，
负责人需先在本机整理、提交并推送这些改动，再通知接手者拉取对应提交。
本次仅按用户要求修改交接形式，没有额外提交或推送代码。

交接附件：
- `WORKTREE_STATUS.txt`：本机基准提交、远程和未提交改动状态。
- `sepolia.json`：当前链上地址、交易及验证清单。

服务器 `/etc/pledge/backend.env` 不进入源码仓库。

## 4. 当前链上部署

网络 Sepolia，chainId `11155111`，poolIndex=0，displayPoolId=1。

| 合约 | 当前地址 |
|---|---|
| PledgePool | `0x1D462BBE7451F555c6742E09972a61AFFfEaD1e9` |
| 六位测试 USDT | `0x9DB37141f5708ADA96AeFE69ce84623880Dfe184` |
| SP | `0x8c2094896f8B2bA54d663CE7Acd1294447e0F164` |
| JP | `0x210bbD576f78da5D1645fcDc8f7325621d89A90f` |
| 多签 | `0xAdC32122226C387CDF0b602bD76Ccc32041eFC8F` |
| 测试价格预言机 | `0xF691AA22Da250f32f55afD139491a0F44C21c259` |
| WETH | `0x841E32459367af4D91C5810dea7979Bf1C42e578` |
| DEX Factory | `0xfE423Fdf5fc9C6df85116F42C036E7e347C5701F` |
| DEX Router | `0xE64974473A2Fc860De6D594d37Ad92846b45CB99` |
| DEX Pair | `0x1bfd0b85c162aa9821d644806398a3d6c3dc0f35` |

部署区块：12055166–12055197。
募集截止：2026-11-04 20:03:00 UTC+8。
到期：2026-11-11 20:03:00 UTC+8。

默认参数：抵押率 200%，年利率 20%，期限 7 天；最低出借金额严格大于 1 USDT；
两端平台费用初始为零；DEX 注入 0.1 测试 ETH + 200 测试 USDT。
手工测试价格为 ETH=2000、USDT=1，原始报价统一为 10^8 精度；不是实时生产价格源。
时间到了不会自动推进状态，目前 Go 同步进程不会发起匹配/结算/清算交易。

主池浏览器：<https://sepolia.etherscan.io/address/0x1D462BBE7451F555c6742E09972a61AFFfEaD1e9>。
完整交易哈希和地址以附件及源码 `pledgev2/networks/sepolia.json` 为准。

### 权限

Owner / 部署执行账户 / 费用接收账户：`0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c`。
多签管理员为三人，门槛 2/3：

- `0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c`
- `0xE7018da7BD70fC8696524257550585f4166E2332`
- `0x6fA81f122444Fe58dcF6388BA2dDE91587565416`

当前 Owner 已获池、Oracle、SP/JP 权限配置的账户级调用授权。
多签审批身份、后台登录身份和凭证 minter 是不同权限；Pool 为凭证 minter。
签署者私钥此前从本机环境读取，未放入后端服务器配置。接手者如需链上操作，
必须由 Owner 提供适当的签名方式，不能假设 SSH 私钥就是链上账户私钥。

## 5. 精度实现要点

- 实际 USDT 金额在池中使用 6 位原始整数，原生 ETH 使用 18 位。
- lendScale[0]=10^12；SP 固定分母为 M_raw × 10^12。
- JP 固定分母为 floor(M_raw × 10^12 × martgageRate / 10^8)。
- 估值先统一到 18 位；可借额向下取整，反算匹配抵押物向上取整。
- 使用 Math.mulDiv；铸造、迟领和赎回使用一致的固定分母，不以动态 totalSupply 分摊。
- 小额 SP 赎回若不足一个 USDT 最小单位，则回滚，保留凭证。
- Oracle 将每整枚资产价格归一化为 10^8，不再用 ERC20 decimals 放大价格；批量价格按资产键写入。
- 只向必要输入资产授权，输出 USDT 不获无限 allowance。
- 当前池合约限定原生 ETH 抵押和同一借款资产；活动借款资产是配置中的六位测试 USDT。
- 所有链上金额通过后端以十进制整数字符串传输，不能用 float64 处理资金。

## 6. 后端现状

目录 `pledge-backend/`，模块名 pledge-backend。Go 要求 1.25+，本次用 1.27.0 验证和构建。
go-ethereum 和 abigen 均为 1.17.8；ABI 由当前 Foundry 编译产物生成。

- API 路由保持原项目命名，前缀 `/api/v1`。
- 同步间隔两分钟，读取最新高度减 12 个确认块；每轮固定区块读取并复核区块哈希。
- 基础数据、分配数据和同步水位在同一 SQL 事务提交，提交成功后才推进 Redis 缓存。
- 同步不会 FLUSHDB；失败不覆盖为零数据；无成功快照、同步失败或超过五分钟时市场接口返回 503。
- 查询按 chainId、poolContract、poolId 隔离；SQL 过滤参数化；分页最多 100 条。
- 后台密码由配置提供，HS256 验证，JWT 与 Redis 会话过期一致；退出使会话失效。
- 只读同步无需链上私钥；默认不启动旧 PLGR 价格订阅、自动喂价或邮件任务。
- 后台 setMultiSign/getMultiSign 只读写后台记录，不完成链上审批。

## 7. 服务器实况与连接

```sh
ssh -i ~/.ssh/t-xjp-server-01 -p 22 ubuntu@43.160.212.151
```

SSH 私钥文件由项目负责人另行安全提供，文档和源码快照不包含它。
服务器 Ubuntu 26.04 LTS / x86_64，约 8GB RAM，2GB Swap，检查时根盘约剩 69GB。
已有 Go 1.26.8、MySQL 8.4.11、Redis 8.0.5、Nginx 1.28.3。

| 组件 | 位置 / 标识 |
|---|---|
| 当前 API、同步二进制 | /opt/pledge/current |
| 当前程序目录 | /opt/pledge/releases/20261103121836-usdt6 |
| 环境配置 | /etc/pledge/backend.env，root:root，0600 |
| MySQL | 127.0.0.1:3306，数据库和账户均为 pledge_study |
| 本项目 Redis | 127.0.0.1:6381，DB 0，独立实例 |
| API | 127.0.0.1:8081 |
| Nginx 路由 | /etc/nginx/snippets/pledge.conf |
| Nginx 配置备份 | /etc/pledge/nginx-backups/ |
| 临时上传目录 | /home/ubuntu/pledge-release，不是源码主目录 |

服务：pledge-api、pledge-sync、pledge-redis，均启用开机自启和失败重启，使用 pledge 系统账户。
服务器还运行其他人的 dbadmin（8080）、原 Redis（6379）、MySQL 其他数据库和 Nginx 路由。
**只维护 Pledge 的命名空间，不删除 app 数据库，不清空原 Redis，不覆盖 /api/、/db/ 路由。**
磁盘上可能仍有历史上传文件或程序目录，它们不是当前功能基线，也不构成存量迁移要求。

### 凭据

后台用户名 admin，密码为随机生成。数据库密码、JWT secret、RPC URL 均在服务器环境文件。
仅在服务器本机或安全会话读取需要的字段，不粘贴到 Git、日志或公开文档。

```sh
sudo awk -F= '/^PLEDGE_ADMIN_(USER|PASSWORD)=/{print}' /etc/pledge/backend.env
```

## 8. 访问与日常检查

公网只读基址：`http://43.160.212.151/pledge`。

```sh
curl -fsS http://43.160.212.151/pledge/healthz
curl -fsS http://43.160.212.151/pledge/api/v1/deployment
curl -fsS http://43.160.212.151/pledge/api/v1/sync/status
curl -fsS 'http://43.160.212.151/pledge/api/v1/poolBaseInfo?chainId=11155111'
curl -fsS 'http://43.160.212.151/pledge/api/v1/token?chainId=11155111'
```

预期：health ready=true；sync complete=true、stale=false；当前池数量 1；
代币精度 ETH=18、USDT=6、SP=18、JP=18。上述状态有时效性，接手时重新检查。

后台 SSH 隧道：

```sh
ssh -i ~/.ssh/t-xjp-server-01 -p 22 -N   -L 18081:127.0.0.1:8081 ubuntu@43.160.212.151
```

隧道接口基址 `http://127.0.0.1:18081/api/v1`。
POST `/user/login` 使用 JSON name/password，随后通过 authCode header 传登录 token。
后台接口包括 `/pool/search`、`/pool/debtTokenList`、`/pool/setMultiSign`、`/pool/getMultiSign`、`/user/logout`。
公网 POST 是 403，这是已选定的无域名访问方案，不是 Go 服务故障。

```sh
sudo systemctl status pledge-api pledge-sync pledge-redis --no-pager
sudo journalctl -u pledge-api -u pledge-sync -n 100 --no-pager
sudo nginx -t
sudo mysql -e 'SELECT pool_contract,pool_id FROM pledge_study.poolbases;'
redis-cli -p 6381 PING
```

健康检查仅代表 DB/Redis 可用，不代表快照新鲜；必须同时检查 sync/status。
同步失败时先查 RPC 网络、部署配置、链上代码和日志；不要把数据库状态手工改成已成交或已结算。

## 9. 构建与验证

从 Git 获取负责人确认的最新提交后，在项目目录执行：

```sh
cd pledgev2
FOUNDRY_PROFILE=sepolia forge test -vv
forge build
cd ../pledge-backend
go test ./...
go build ./...
go vet ./...
make chaincheck  # 需提供 SEPOLIA_RPC_URL，不需要数据库或链上私钥
```

Foundry 编译器 0.8.24、OpenZeppelin 4.9.6、forge-std 1.9.7，依赖已随源码放在 lib。
Sepolia 运行时使用 Amsterdam；编译目标 Paris。
曾因旧 Gas 模型低估导致初始部署失败，已修正；不要用旧 Gas 估算重试当前部署。
`make deploy-sepolia` 会创建新部署，不能用于日常启动或只补源码验证。
源码补验可使用 `tools/verify-exact-build.py`，需要准确构建信息及 Etherscan API key。

重新生成 Go 绑定：先在 `pledgev2/` 执行 `forge build`，再在后端 `make bindings`；
这一步也会同步 config/sepolia.json。配置通过 go:embed 编入二进制，改变嵌入地址后需重新构建。
服务器 EnvironmentFile 中显式配置的变量按现有权限管理；后端源码 README 提供完整变量说明。

本次测试证据：
- Foundry 26 个用例通过；两组模糊测试各 256 轮。
- Go 三个集成回归用例通过；测试内使用 SQLite + 进程内 Redis。
- 实际服务器 MySQL/Redis 初始化、登录、鉴权搜索、退出、链上同步及公开查询通过。
- 当前 31 笔链上部署交易和 10 份代码/源码已核对。
- Docker/Compose 构建没有实测；实际服务器使用 systemd，不使用容器。

## 10. 仍未完成的工作（优先接续）

### 合约 P0

- 清算入口没有强制风险阈值，授权账户可在查询条件不成立时清算。
- 提前清算仍按完整约定期限计息；不足额抵押物处置、损失分摊和终态退出尚未完整解决。
- 兑换最小产出仍为零，deadline 在执行时生成；缺少调用级最大支出、最小产出和报价有效期。
- 多签仍是账户+目标的可复用授权，无逐笔函数/参数/nonce/到期绑定。
- 全局费率和路由变更可影响存续池，未冻结费用快照。
- 池级预算隔离、跨池凭证复用限制、舍入余量归属等仍需完善。
- ETH transfer 对部分合约钱包不兼容；价格轮次基本校验不等于最大时效保护。
- 先明确 PRD 决策再修改对应资金规则；通过精度测试不代表上述问题已完成。

### 后端与前端

- 用户仓位、纯凭证受让人发现、可领金额和历史查询。
- 完整事件索引、原生 ETH 支付核对、交易确认和重组回滚。
- 匹配/结算/清算执行队列、授权签名、幂等、重试和人工接管；当前只有读链同步。
- 多用户角色、管理操作审计、历史配置版本和资金对账。
- 性能、告警和端到端验收；前端尚未复制建设。

建议接续顺序：先整理、提交并推送当前代码 → 复核 PRD P0 决策 → 补合约资金保护与测试 →
补仓位/事件索引和执行服务 → 建设前端 → 测试网端到端验收。

## 11. 资料入口

优先看 `Pledge_PRD.md` 第 17.2.3 节（已确认精度规则）、`docs/usdt6-migration.md`、
`docs/sepolia-deployment.md`、`docs/backend-server-deployment.md`、`pledge-backend/README.md`。
部分早期实施记录描述历史基线，若与当前六位 USDT 部署不同，以本文件、当前代码和链上记录为准。

当前服务二进制 SHA-256：

```text
pledge-api   8bd77799515f4a00aecaae3a36d4603ec87e02a0205a0cdcfa31b0aafdd62ac2
pledge-sync  9dd6b0647d09f014549fdb93cf7caceb1c0364dea229585f977976e072943579
```
