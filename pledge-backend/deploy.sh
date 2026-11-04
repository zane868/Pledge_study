#!/bin/bash
# Pledge 后端部署脚本
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
