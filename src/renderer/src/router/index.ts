import { createRouter, createWebHashHistory } from 'vue-router'
import { constantRouterMap } from '@renderer/config/router.config'
import { UserAuthStore } from '@renderer/store/user'

const router = createRouter({
  // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHashHistory(),
  routes: constantRouterMap // `routes: routes` 的缩写
})

// GOOD
router.beforeEach((to, _from, next) => {
  const user = UserAuthStore()
  if (to.meta.requireAuth) {
    if (user.isAuthenticated) {
      next()
    } else {
      next({ path: '/user' })
    }
  } else {
    next()
  }
})

export default router
