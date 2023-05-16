import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'),
            'components': path.resolve(__dirname, 'src/components'),
            'views': path.resolve(__dirname, 'src/views'),
            'utils': path.resolve(__dirname, 'src/utils'),
            'assets': path.resolve(__dirname, 'src/assets'),
            'api': path.resolve(__dirname, 'src/api'),
        }
    },
    define: {
        'process.env': process.env
    },
    build: {
        target: 'modules',
        outDir: '../dist', //指定输出路径
        assetsDir: 'assets', // 指定生成静态资源的存放路径
        // minify: 'terser' // 混淆器，terser构建后文件体积更小
    },
    server: {
        cors: true, // 默认启用并允许任何源
        open: true, // 在服务器启动时自动在浏览器中打开应用程序
        port: 8080,
        //反向代理配置，注意rewrite写法
        proxy: {
            '^/api': {
                target: 'https://127.0.0.1:5000',   //代理接口
                changeOrigin: true,
                secure: false
                // rewrite: (path) => path.replace(/^\/api/, '/v1')
            },
            '^/uploads': {
                target: 'https://127.0.0.1:5000',   //代理资源
                changeOrigin: true,
                secure: false
                // rewrite: (path) => path.replace(/^\/api/, '/v1')
            }
        }
    }
})
