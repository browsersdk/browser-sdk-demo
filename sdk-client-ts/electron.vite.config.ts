import { resolve } from 'path'
import { defineConfig } from 'electron-vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  main: {
    resolve: {
      alias: {
        '@src2': resolve(__dirname, 'src/*'),
        '@main': resolve(__dirname, 'src/main'),
        '@type': resolve(__dirname, 'src/types')
      }
    },
    build: {
      rollupOptions: {
        input: {
          index: 'src/main/index.ts'
        },
        output: {
          entryFileNames: '[name].js'
          // format: 'cjs'
        },
        external: ['koffi', 'electron-log']
      }
    }
  },
  preload: {
    build: {
      rollupOptions: {
        external: ['@electron-toolkit/preload']
      }
    }
  },
  renderer: {
    resolve: {
      alias: {
        '@': resolve('src/renderer/src'),
        '@src': resolve(__dirname, 'src'),
        '@type': resolve('src/types')
      }
    },
    plugins: [vue()]
  }
})
