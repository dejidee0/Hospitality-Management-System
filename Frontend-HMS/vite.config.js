import { defineConfig } from 'vite'
import path from 'path';
import react from '@vitejs/plugin-react'
import svgr from 'vite-plugin-svgr';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), svgr()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@components': path.resolve(__dirname, 'src/components'),
      '@features': path.resolve(__dirname, 'src/features'),
      '@api/*': path.resolve(__dirname, 'src/api'),
      '@hooks/*': path.resolve(__dirname, 'src/hooks'),
      '@redux/*': path.resolve(__dirname, 'src/redux'),
      '@pages/*': path.resolve(__dirname, 'src/pages'),
      '@routes/*': path.resolve(__dirname, 'src/routes'),
      '@assets/*': path.resolve(__dirname, 'src/assets'),
      '@utils/*': path.resolve(__dirname, 'src/utils'),
    },
  },
});
