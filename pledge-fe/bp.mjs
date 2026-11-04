import { build } from 'vite';
import react from '@vitejs/plugin-react';
const r = await build({
  plugins: [react()],
  logLevel: 'silent',
  configFile: false,
  build: { write: false, minify: false, lib: false },
  root: '.',
});
