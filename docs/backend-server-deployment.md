# Go 后端服务器部署

部署日期：2026-11-03。服务器：`ubuntu@43.160.212.151:22`。

## 环境检查

Ubuntu 26.04 LTS，x86_64，约 8GB 内存、2GB Swap、根盘剩余 69GB。
已安装 Go 1.26.8、MySQL 8.4.11、Redis 8.0.5、Nginx 1.28.3，无需新增 Docker。
原有 dbadmin 使用 8080，MySQL/Redis 已有其他应用使用。

本项目使用独立 MySQL 数据库及账户 `pledge_study`，独立 Redis 实例监听
`127.0.0.1:6381`，Go API 监听 `127.0.0.1:8081`。原应用及数据库未被替换或清空。

## 访问

- 健康检查：<http://43.160.212.151/pledge/healthz>
- 池信息：<http://43.160.212.151/pledge/api/v1/poolBaseInfo?chainId=11155111>
- 分配数据：<http://43.160.212.151/pledge/api/v1/poolDataInfo?chainId=11155111>
- 同步状态：<http://43.160.212.151/pledge/api/v1/sync/status>
- 部署元数据：<http://43.160.212.151/pledge/api/v1/deployment>

用户确认无需域名，按无域名方案通过 Nginx `/pledge/` 转发只读请求。
公网 POST 返回 403；需要登录的后台接口通过 SSH 隧道访问：

```sh
ssh -i ~/.ssh/t-xjp-server-01 -p 22 -N \
  -L 18081:127.0.0.1:8081 ubuntu@43.160.212.151
```

隧道地址：`http://127.0.0.1:18081/api/v1`。
后台用户名为 `admin`，随机密码和其他连接凭据仅保存在服务器
`/etc/pledge/backend.env`（root 所有、权限 0600），没有写入仓库。
登录后用 `authCode` header 调用后台接口。

仅查看后台登录凭据：

```sh
sudo awk -F= '/^PLEDGE_ADMIN_(USER|PASSWORD)=/{print}' /etc/pledge/backend.env
```

## 服务与维护

三个服务均已配置开机启动和失败重启：

```sh
sudo systemctl status pledge-api pledge-sync pledge-redis
sudo journalctl -u pledge-api -u pledge-sync -n 100 --no-pager
```

程序目录：`/opt/pledge/current`，当前指向 `/opt/pledge/releases/20261103121836-usdt6`。
项目只使用当前 USDT6 部署。
配置：`/etc/pledge/backend.env`；Nginx 路由：`/etc/nginx/snippets/pledge.conf`。
原 Nginx 配置备份位于 `/etc/pledge/nginx-backups/`。
首次安装脚本保存在 `pledge-backend/ops/`，会拒绝覆盖已有后端配置，不用于直接重复部署。

本次程序使用 Go 1.27.0 交叉编译为 Linux amd64 静态二进制。
## 验证

已通过服务器实际 MySQL/Redis 初始化、API 健康检查、Sepolia 同步、公开池查询、
后台登录、鉴权搜索和退出验证。同步状态返回完整快照，池数量为 1，且未过期。
原 dbadmin、Nginx、MySQL 及原 Redis 服务仍正常运行。

## USDT 六位精度更新

新版活动池为 `0x1D462BBE7451F555c6742E09972a61AFFfEaD1e9`，测试 USDT 为
`0x9DB37141f5708ADA96AeFE69ce84623880Dfe184`。公开接口已确认 ETH 18 位、USDT 6 位、SP/JP 18 位。
新快照同步成功，项目只保留活动池测试记录；数据库密码、后台登录密码及 JWT secret 沿用原配置。
当前程序 SHA-256：

```text
pledge-api   8bd77799515f4a00aecaae3a36d4603ec87e02a0205a0cdcfa31b0aafdd62ac2
pledge-sync  9dd6b0647d09f014549fdb93cf7caceb1c0364dea229585f977976e072943579
```

后续按全新系统维护，不再提供旧资产迁移或旧版本切换流程。
