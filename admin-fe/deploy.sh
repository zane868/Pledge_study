#!/bin/bash
# Pledge 运营后台部署脚本
# 服务器: ubuntu@43.160.212.151

set -e

SERVER="ubuntu@43.160.212.151"
SSH_KEY="$HOME/.ssh/t-xjp-server-01"
REMOTE_DIR="/var/www/pledge-admin"
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
echo "=== 4. 配置 Nginx（首次部署需手动执行） ==="
echo "首次部署请在服务器上执行："
echo "  sudo cp /tmp/nginx-admin.conf /etc/nginx/snippets/pledge-admin.conf"
echo "  sudo vi /etc/nginx/sites-enabled/default  # 添加 include snippets/pledge-admin.conf;"
echo "  sudo nginx -t && sudo systemctl reload nginx"
echo ""
echo "上传 Nginx 配置到服务器..."
scp -i $SSH_KEY nginx-admin.conf $SERVER:/tmp/nginx-admin.conf

echo ""
echo "=== 部署完成 ==="
echo "访问地址: http://43.160.212.151/admin/"
echo ""
echo "清理旧版本（保留最近 3 个）："
echo "  ssh -i $SSH_KEY $SERVER 'cd $REMOTE_DIR/releases && ls -t | tail -n +4 | xargs rm -rf'"
