# 前端服务器部署

部署日期：2026-11-04。服务器：`ubuntu@43.160.212.151:22`。

## 访问

- 前端：<http://43.160.212.151/app/>
- 后端只读接口：<http://43.160.212.151/pledge/api/v1/...>
- 健康检查：<http://43.160.212.151/pledge/healthz>

前端是纯静态单页应用，部署在 Nginx 的 `/app/` 子路径下，与根路径的演示页、
`/db/`（数据库管理）、`/redis/`（Redis 管理）、`/pledge/`（后端 API）互不干扰。

## 目录结构

```
/var/www/pledge-fe/
├── current -> releases/20261104140200   # 切换版本只改这个软链
└── releases/
    ├── 20261104134645/
    └── 20261104140200/
        ├── index.html
        └── assets/                      # 带内容哈希，可长期缓存
```

Nginx 配置：`/etc/nginx/snippets/pledge-fe.conf`，在
`/etc/nginx/sites-enabled/default` 的 server 块中 include。
改动前已备份为 `default.bak.<时间戳>`。

## 构建与发布

```sh
cd pledge-fe
npm ci
npm run build          # 产物在 pledge-fe/dist
```

`vite.config.ts` 固定 `base: '/app/'`，`main.tsx` 的 `BrowserRouter`
用 `import.meta.env.BASE_URL` 作为 basename —— 两者必须同时是 `/app/`，
否则资源 404 或前端路由失效。

发布到服务器：

```sh
REL=$(date +%Y%m%d%H%M%S)
COPYFILE_DISABLE=1 tar -czf /tmp/pledge-fe.tar.gz -C pledge-fe/dist .
scp -i ~/.ssh/t-xjp-server-01 /tmp/pledge-fe.tar.gz ubuntu@43.160.212.151:/tmp/

ssh -i ~/.ssh/t-xjp-server-01 ubuntu@43.160.212.151 "
  sudo mkdir -p /var/www/pledge-fe/releases/$REL &&
  sudo chown ubuntu:ubuntu /var/www/pledge-fe/releases/$REL &&
  tar -xzf /tmp/pledge-fe.tar.gz -C /var/www/pledge-fe/releases/$REL &&
  sudo ln -sfn /var/www/pledge-fe/releases/$REL /var/www/pledge-fe/current
"
```

`COPYFILE_DISABLE=1` 用于避免 macOS 在 tar 包中附加 `._*` 资源分叉文件。

## 回滚

切回上一个版本只需重指软链：

```sh
sudo ln -sfn /var/www/pledge-fe/releases/<上一个版本> /var/www/pledge-fe/current
```

彻底移除：

```sh
sudo rm /etc/nginx/snippets/pledge-fe.conf
sudo sed -i '/pledge-fe.conf/d' /etc/nginx/sites-enabled/default
sudo nginx -t && sudo systemctl reload nginx
sudo rm -rf /var/www/pledge-fe
```

## Nginx 路由要点

```nginx
location = /app { return 301 /app/; }

location ^~ /app/ {
    alias /var/www/pledge-fe/current/;
    try_files $uri $uri/ /app/index.html;   # SPA 回退
}

location ^~ /app/assets/ {
    alias /var/www/pledge-fe/current/assets/;
    expires 30d;
    add_header Cache-Control "public, immutable";
}
```

`try_files` 中的回退目标必须是外部 URI（`/app/index.html`），不能写成
`/index.html` —— 后者会落到根路径的演示页。

## 网络与数据来源

- 链上读取与交易：浏览器直接用 `shared/deployment.json` 里的 Sepolia RPC，
  不经过服务器；交易由用户钱包签名后直接发往合约。
- 市场列表：调用同域 `/pledge/api/v1/poolBaseInfo`，由现有 Go 后端提供。
- 后端在同步未就绪或快照超过 5 分钟时返回 503，前端区分「加载失败」与
  「确实没有池」，不会把失败显示成空市场。

## 验证

```sh
for p in /app/ /app/portfolio /pledge/healthz /db/ /; do
  printf '%-20s ' "$p"; curl -s -o /dev/null -w '%{http_code}\n' "http://43.160.212.151$p"
done
```

预期全部 200。