import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'highlight.js/lib/common';
import 'highlight.js/styles/atom-one-dark-reasonable.css'
import HljsPlugin from '@highlightjs/vue-plugin'
import zh_local from 'element-plus/dist/locale/zh-cn'
import App from './App.vue'
import router from './router'

import './assets/main.css'

let lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string

const app = createApp(App)

app.use(createPinia())
app.use(router)
if (lang == 'zh') {
    app.use(ElementPlus, { locale: zh_local })
} else {
    app.use(ElementPlus)
}

app.use(HljsPlugin)

app.mount('#app')
