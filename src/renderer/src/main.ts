import 'ant-design-vue/dist/antd.css'
import '@ant-design-vue/pro-layout/dist/style.css'
import 'ant-design-vue/dist/antd.variable.min.css'
import '@renderer/global.less'

import { createApp } from 'vue'
import App from '@renderer/App.vue'
import router from '@renderer/router'
import { createPinia } from 'pinia'
import commonFunc from '@renderer/util/commonFunc'
import ProLayout, { PageContainer } from '@ant-design-vue/pro-layout'

const app = createApp(App)

app
  .use(router)
  .use(commonFunc)
  .use(createPinia())
  .use(ProLayout)
  .use(PageContainer)
  //将这个vue3app全局挂载到#app元素上
  .mount('#app')
