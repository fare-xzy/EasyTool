import request from '@renderer/util/request'
import { notification, message } from 'ant-design-vue'

const $timeFix = (): string => {
  const time = new Date()
  const hour = time.getHours()
  return hour < 9
    ? '早上好'
    : hour <= 11
    ? '上午好'
    : hour <= 13
    ? '中午好'
    : hour < 20
    ? '下午好'
    : '晚上好'
}

export default {
  install: (app) => {
    app.config.globalProperties['$timeFix'] = $timeFix
    app.config.globalProperties.$axios = request
    app.config.globalProperties.$notification = notification
    app.config.globalProperties.$message = message
  }
}
