import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  optimizeDeps: {
    exclude: ['lucide-react'],
  },
  server: {
    // allowedHosts:true,
    proxy: {
      '/api': {
        // target: 'http://localhost:5000',
        target: 'https://hassisconecta.com.br',
        changeOrigin: true,
        secure: false,
      },
      '/images/upload': {
        // target: 'http://localhost:5000',
        target: 'https://hassisconecta.com.br',
        changeOrigin: true,
        secure: false,
      },
    },
  }
});
