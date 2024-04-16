import react from '@vitejs/plugin-react'
import path from 'path'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
const currentDir = path.dirname(import.meta.url)

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@data': currentDir+ '/../result/result.json'
    }
  }
})
