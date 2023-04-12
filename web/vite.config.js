import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@': '/src',
            'assets': '@/assets',
            'components': '@/components',
            'views': '@/views',
        }
    },
    define: {
        'process.env': process.env
    },
    build: {
        target: 'modules',
        outDir: '../build/dist', //指定输出路径
        assetsDir: 'assets', // 指定生成静态资源的存放路径
        // minify: 'terser' // 混淆器，terser构建后文件体积更小
    },
    server: {
        cors: true, // 默认启用并允许任何源
        open: true, // 在服务器启动时自动在浏览器中打开应用程序
        port: 8080,
        //反向代理配置，注意rewrite写法，开始没看文档在这里踩了坑
        proxy: {
            '^/api': {
                target: 'http://127.0.0.1',   //代理接口
                changeOrigin: true,
                // rewrite: (path) => path.replace(/^\/api/, '/v1')
            }
        }
    }
})
