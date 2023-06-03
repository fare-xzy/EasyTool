import { defineStore } from 'pinia'
import { login } from '@renderer/api/user'
import { ReturnResult } from '@renderer/util/request'

export const UserAuthStore = defineStore('UserAuth', {
  state: () => ({
    isAuthenticated: false
  }),
  actions: {
    Login(userInfo): Promise<ReturnResult<object>> {
      return new Promise((resolve, reject) => {
        login(userInfo)
          .then((response) => {
            resolve(response.data)
          })
          .catch((error) => {
            reject(error)
          })
      })
    }
  }
})
