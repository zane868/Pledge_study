# Pledge 项目问题排查手册

## 目录
- [重新部署合约后前端显示旧数据](#重新部署合约后前端显示旧数据)
- [常见问题速查表](#常见问题速查表)

---

## 重新部署合约后前端显示旧数据

### 问题现象
重新部署智能合约后，前端仍然显示旧数据（如旧的出借金额、错误的池子状态等），即使强制刷新浏览器也无效。

### 问题根因
前端显示的数据来自后端 API，后端同步服务使用**编译时嵌入**的配置文件（`//go:embed sepolia.json`）。即使更新了配置文件，旧的二进制文件仍然包含旧的合约地址。

### 排查步骤

#### 1. 确认链上数据是否正确

```bash
export RPC=https://sepolia.gateway.tenderly.co
export POOL=0x7F446d239BA877c4190Fd1394bD2E67840E93345  # 新部署的地址

# 检查池子状态
cast call $POOL "getPoolState(uint256)(uint256)" 0 --rpc-url $RPC

# 检查池子基础信息
cast call $POOL "poolBaseInfo(uint256)" 0 --rpc-url $RPC
```

**预期结果**：链上数据应该是全新的（lendSupply=0, borrowSupply=0）

---

#### 2. 检查前端配置是否更新

```bash
cd /Users/jiandong/github_space/Pledge_study

# 检查 shared/deployment.json 中的地址
cat shared/deployment.json | jq '.pool'

# 如果地址不正确，手动更新或从合约部署文件同步
cat pledgev2/deployments/sepolia.json | jq '{
  pool, oracle, multisig, stable, weth, router, factory, sp, jp,
  deploymentStartBlock: .firstBlock
}'
```

**更新前端配置后，重新部署前端：**

```bash
# 管理后台
cd admin-fe
bash deploy.sh

# 对客前端
cd ../pledge-fe
bash deploy.sh
```

---

#### 3. 检查后端 API 返回的地址（关键步骤）

```bash
# 检查后端同步服务返回的合约地址
curl -s 'http://43.160.212.151/pledge/api/v1/sync/status' | jq '.data.poolContract'
```

**如果返回的是旧地址** → 后端配置未更新，继续下一步

**如果返回的是新地址** → 后端正常，检查前端缓存问题

---

#### 4. 更新后端配置并重新部署

##### 4.1 更新后端配置文件

```bash
cd /Users/jiandong/github_space/Pledge_study/pledge-backend

# 编辑 config/sepolia.json，更新所有合约地址
# 从 pledgev2/deployments/sepolia.json 复制新地址
cat ../pledgev2/deployments/sepolia.json
```

**必须更新的字段：**
- `pool` - PledgePool 合约地址
- `oracle` - BscPledgeOracle 地址
- `multisig` - 多签合约地址
- `stable` - 稳定币地址（TestUSDT）
- `weth` - WETH 地址
- `router` - UniswapV2Router02 地址
- `factory` - UniswapV2Factory 地址
- `sp` - Senior Pool Token 地址
- `jp` - Junior Pool Token 地址
- `deploymentStartBlock` - 新部署的起始区块号

##### 4.2 重新编译并部署后端

**重要：** 因为配置是通过 `//go:embed` 编译时嵌入的，**必须重新编译**二进制文件。

```bash
cd /Users/jiandong/github_space/Pledge_study/pledge-backend

# 如果没有 deploy.sh，创建部署脚本
cat > deploy.sh << 'EOF'
#!/bin/bash
set -e

SERVER="ubuntu@43.160.212.151"
SSH_KEY="$HOME/.ssh/t-xjp-server-01"
REMOTE_DIR="/opt/pledge"

echo "=== 1. 编译后端（交叉编译 Linux amd64）==="
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pledge-api ./api
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pledge-sync ./schedule

echo ""
echo "=== 2. 上传二进制文件 ==="
scp -i $SSH_KEY bin/pledge-api bin/pledge-sync $SERVER:/tmp/

echo ""
echo "=== 3. 替换并重启服务 ==="
ssh -i $SSH_KEY $SERVER "sudo mv /tmp/pledge-api /tmp/pledge-sync $REMOTE_DIR/current/ && \
  sudo chown root:root $REMOTE_DIR/current/pledge-* && \
  sudo chmod +x $REMOTE_DIR/current/pledge-* && \
  sudo systemctl restart pledge-api pledge-sync && \
  sleep 3 && \
  sudo systemctl status pledge-api pledge-sync --no-pager"

echo ""
echo "=== 部署完成 ==="
EOF

chmod +x deploy.sh

# 执行部署
bash deploy.sh
```

**为什么需要交叉编译？**
- 本地开发环境：macOS ARM64 (M1/M2/M3)
- 服务器环境：Linux x86_64
- 必须指定 `GOOS=linux GOARCH=amd64`

---

#### 5. 验证修复结果

```bash
# 等待服务重启并同步（约 5-10 秒）
sleep 10

# 再次检查 API 返回的合约地址
curl -s 'http://43.160.212.151/pledge/api/v1/sync/status' | jq '{
  poolContract: .data.poolContract,
  blockNumber: .data.blockNumber,
  lastSuccess: .data.lastSuccess
}'
```

**预期结果：**
```json
{
  "poolContract": "0x7F446d239BA877c4190Fd1394bD2E67840E93345",  // 新地址
  "blockNumber": "12062544",
  "lastSuccess": "2026-11-04T12:56:42Z"
}
```

---

#### 6. 清除前端缓存并验证

- **强制刷新浏览器：**
  - Windows/Linux: `Ctrl + Shift + R`
  - macOS: `Cmd + Shift + R`
  
- **或使用开发者工具：**
  1. 按 `F12` 打开开发者工具
  2. 切换到 `Network` 标签
  3. 勾选 `Disable cache`
  4. 刷新页面

**验证前端显示：**
- 管理后台：http://43.160.212.151/admin/
- 对客前端：http://43.160.212.151/app/

**应该看到：**
- ✅ 池子状态显示 "募集中"（如果当前时间 < settleTime）
- ✅ 出借金额显示 0
- ✅ 借款金额显示 0
- ✅ 合约地址已更新

---

## 常见问题速查表

| 问题 | 可能原因 | 快速检查命令 | 解决方案 |
|------|---------|------------|---------|
| 前端显示旧数据 | 后端 API 返回旧地址 | `curl 'http://43.160.212.151/pledge/api/v1/sync/status' \| jq '.data.poolContract'` | 更新后端配置并重新编译部署 |
| 后端服务启动失败 (status=203) | 二进制文件架构不匹配 | `ssh user@server 'file /opt/pledge/current/pledge-api'` | 使用交叉编译：`GOOS=linux GOARCH=amd64 go build` |
| 修改配置后重启无效 | 配置是编译时嵌入的 | `grep -rn "go:embed" config/` | 必须重新编译二进制文件 |
| 池子显示"募集已截止" | settleTime 已过期 | `cast call $POOL "poolBaseInfo(uint256)" 0` | 通过管理后台创建新池子或修改时间 |
| 浏览器缓存未清除 | 前端 JS 文件被缓存 | 开发者工具 Network 查看请求 | `Ctrl+Shift+R` 强制刷新或禁用缓存 |

---

## 部署流程总结

### 完整的合约重新部署流程

```bash
# 1. 重新部署合约
cd pledgev2
forge script script/DeploySepolia.s.sol --rpc-url $SEPOLIA_RPC_URL --broadcast

# 2. 更新 shared/deployment.json（前端共享配置）
cat deployments/sepolia.json | jq '{
  chainId, pool, oracle, multisig, stable, weth, router, factory, sp, jp,
  operator, admins, threshold, stableDecimals,
  deploymentStartBlock: .firstBlock
}' > ../shared/deployment.json

# 3. 更新后端配置
cp deployments/sepolia.json ../pledge-backend/config/sepolia.json

# 4. 重新编译并部署后端
cd ../pledge-backend
bash deploy.sh

# 5. 重新部署前端
cd ../admin-fe
bash deploy.sh

cd ../pledge-fe
bash deploy.sh

# 6. 验证
curl -s 'http://43.160.212.151/pledge/api/v1/sync/status' | jq '.data.poolContract'
```

---

## 调试技巧

### 查看后端日志

```bash
# 查看 API 服务日志
ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 \
  "sudo journalctl -u pledge-api -n 50 --no-pager"

# 查看同步服务日志
ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 \
  "sudo journalctl -u pledge-sync -n 50 --no-pager"

# 实时跟踪日志
ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 \
  "sudo journalctl -u pledge-sync -f"
```

### 检查服务状态

```bash
ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 \
  "sudo systemctl status pledge-api pledge-sync --no-pager"
```

### 手动重启服务

```bash
ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 \
  "sudo systemctl restart pledge-api pledge-sync"
```

---

## 关键要点总结

1. **后端配置是编译时嵌入的（`//go:embed`）**
   - 修改配置文件后必须重新编译二进制文件
   - 仅重启服务不会加载新配置

2. **交叉编译时注意目标平台架构**
   - macOS → Linux 必须指定 `GOOS=linux GOARCH=amd64`
   - 否则会出现 status=203 启动失败

3. **前端、后端、链上数据必须同步更新**
   - 合约地址变更时，三者必须一起更新
   - 使用 `shared/deployment.json` 作为单一配置源

4. **验证时检查后端 API 返回的地址**
   - 不要只看配置文件
   - API 返回的地址才是实际使用的地址

5. **浏览器缓存问题**
   - 部署后必须强制刷新（`Ctrl+Shift+R`）
   - 开发时可开启 "Disable cache"

---

## 相关文件路径

- 合约部署输出：`pledgev2/deployments/sepolia.json`
- 前端共享配置：`shared/deployment.json`
- 后端配置文件：`pledge-backend/config/sepolia.json`
- 管理后台部署脚本：`admin-fe/deploy.sh`
- 对客前端部署脚本：`pledge-fe/deploy.sh`
- 后端部署脚本：`pledge-backend/deploy.sh`

---

## 联系与支持

遇到其他问题请查看：
- 合约文档：`pledgev2/README.md`
- 后端文档：`pledge-backend/README.md`
- 前端文档：`admin-fe/README.md`, `pledge-fe/README.md`
