import { constantRouterMap } from '@renderer/config/router.config'
import { defineStore } from 'pinia'

export const PermissionPropsStore = defineStore('PermissionProps', {
  state: () => ({
    routers: constantRouterMap
  })
})
