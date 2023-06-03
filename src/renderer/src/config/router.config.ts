import { UserLayout, BasicLayout } from '@renderer/layouts'
import { Login, TableList, History } from '@renderer/views/'
import { IconUpdate } from '@renderer/assets/icons'

const RouteView = {
  name: 'RouteView',
  // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
  render: () => h(resolveComponent('router-view'))
}

export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/login',
    children: [
      {
        path: '/login',
        name: 'login',
        component: Login
      }
    ]
  },
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { requireAuth: true },
    redirect: '/list/table-list',
    children: [
      // dashboard
      {
        path: '/list',
        name: 'list',
        component: RouteView,
        redirect: '/list/table-list',
        meta: { title: '升级模快', icon: IconUpdate, requireAuth: true },
        children: [
          {
            path: '/list/table-list/:pageNo([1-9]\\d*)?',
            name: 'TableListWrapper',
            hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
            component: TableList,
            meta: { title: '升级', keepAlive: true, requireAuth: true }
          },
          {
            path: '/list/basic-list',
            name: 'BasicList',
            component: History,
            meta: { title: '升级历史', keepAlive: true, requireAuth: true }
          }
        ]
      }
    ]
  }
]
