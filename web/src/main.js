import {createApp} from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
import Vant from 'vant'
import {Lazyload} from 'vant';
import 'vant/lib/index.css'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import VueClipBoard from 'vue-clipboard2'

import App from './App.vue'
import router from './router'
import store from './store'


const app = createApp(App)

app.use(store)
app.use(router)

app.use(VueClipBoard)
app.use(ElementPlus, {
    locale: zhCn,
})
app.use(Vant)
app.use(Lazyload)
app.mount('#app')
