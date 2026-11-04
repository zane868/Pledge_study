import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'node:path';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@shared': path.resolve(__dirname, '../shared/src'),
      '@shared-abi': path.resolve(__dirname, '../shared/abi'),
      '@': path.resolve(__dirname, 'src'),
    },
  },
  // 部署在 Nginx 的 /app/ 子路径下，资源引用需带该前缀
  base: '/app/',
  server: {
    port: 3000,
    proxy: {
      // 后端公网只读接口；后台接口走 SSH 隧道，不在此代理
      '/pledge': {
        target: 'http://43.160.212.151',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
  },
});