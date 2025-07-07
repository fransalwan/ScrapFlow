// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' // ⬅️ Tambahkan ini

export default defineConfig({
  plugins: [vue()],
   resolve: {
    alias: {
      '@': path.resolve(__dirname, './src') // ⬅️ Alias ke folder src
    }
  }
})
