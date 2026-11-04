#!/bin/bash
# Pledge 对客前端部署脚本
# 服务器: ubuntu@43.160.212.151

set -e

SERVER="ubuntu@43.160.212.151"
SSH_KEY="$HOME/.ssh/t-xjp-server-01"
REMOTE_DIR="/var/www/pledge-fe"
RELEASE_DIR="$REMOTE_DIR/releases/$(date +%Y%m%d%H%M%S)"

echo "=== 1. 构建前端 ==="
npm run build

echo ""
echo "=== 2. 上传构建产物 ==="
ssh -i $SSH_KEY $SERVER "mkdir -p $RELEASE_DIR"
scp -i $SSH_KEY -r dist/* $SERVER:$RELEASE_DIR/

echo ""
echo "=== 3. 切换当前版本 ==="
ssh -i $SSH_KEY $SERVER "ln -sfn $RELEASE_DIR $REMOTE_DIR/current"

echo ""
echo "=== 部署完成 ==="
echo "访问地址: http://43.160.212.151/app/"
echo ""
echo "清理旧版本（保留最近 3 个）："
echo "  ssh -i $SSH_KEY $SERVER 'cd $REMOTE_DIR/releases && ls -t | tail -n +4 | xargs rm -rf'"
