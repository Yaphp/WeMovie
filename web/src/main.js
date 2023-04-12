import {createApp} from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import BaiduMap from 'vue-baidu-map-3x';

import App from './App.vue'
import router from './router'
import store from './store'

const app = createApp(App)

app.use(store).use(router)
    .use(ElementPlus, {
        locale: zhCn,
    }).use(BaiduMap, {
    ak: 'xmzV94OpCNnSPWDBSBsOya7h6iPCzzYC'
}).mount('#app')
