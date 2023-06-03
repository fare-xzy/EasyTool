import { defineStore } from 'pinia'

export const AppPropsStore = defineStore('AppProps', {
  state: () => ({
    sideCollapsed: false,
    isMobile: false,
    theme: 'dark',
    layout: '',
    contentWidth: '',
    fixedHeader: false,
    fixedSidebar: false,
    autoHideHeader: false,
    color: '',
    weak: false,
    multiTab: true,
    lang: 'en-US',
    _antLocale: {}
  })
})
