<template>
  <pro-layout
    v-model:collapsed="baseState.collapsed"
    v-model:selectedKeys="baseState.selectedKeys"
    v-model:openKeys="baseState.openKeys"
    :menu-data="menuData"
    :breadcrumb="{ routes: breadcrumb }"
    v-bind="settings"
  >
    <template #menuHeaderRender>
      <router-link :to="{ path: '/' }">
        <img src="../assets/logo.jpg" class="logo" alt="logo" />
        <h1>{{ defaultSettings.title }}</h1>
      </router-link>
    </template>

    <!-- custom right-content -->
    <template #rightContentRender>
      <RightContent :current-user="currentUser" />
    </template>
    <!-- <template #rightContentRender>
      <div style="margin-right: 12px">
        <a-avatar shape="square">
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
      </div>
    </template> -->
    <!-- custom breadcrumb itemRender  -->
    <template #breadcrumbRender="{ route, params, routes }">
      <span v-if="routes.indexOf(route) === routes.length - 1">{{ route.breadcrumbName }}</span>
      <RouterLink v-else :to="{ path: route.path, params }">
        {{ route.breadcrumbName }}
      </RouterLink>
    </template>
    <!-- content begin -->
    <RouterView v-slot="{ Component }">
      <!-- <WaterMark :content="watermarkContent"> -->
      <component :is="Component" />
      <!-- </WaterMark> -->
    </RouterView>
  </pro-layout>
</template>

<script setup lang="ts">
import { computed, reactive, ref, watchEffect, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getMenuData, clearMenuItem, type RouteContextProps } from '@ant-design-vue/pro-layout'
import defaultSettings from '@renderer/config/defaultSettings'
import { CONTENT_WIDTH_TYPE } from '@renderer/store/mutation-types'

const watermarkContent = ref('Pro Layout')
const router = useRouter()
const { menuData } = getMenuData(clearMenuItem(router.getRoutes()))
const baseState = reactive<Omit<RouteContextProps, 'menuData'>>({
  selectedKeys: [],
  openKeys: [],
  // default
  collapsed: false
})

const settings = reactive({
  // 布局类型
  layout: defaultSettings.layout, // 'sidemenu', 'topmenu'
  // CONTENT_WIDTH_TYPE
  contentWidth:
    defaultSettings.layout === 'sidemenu' ? CONTENT_WIDTH_TYPE.Fluid : defaultSettings.contentWidth,
  // 主题 'dark' | 'light'
  theme: defaultSettings.navTheme,
  // 主色调
  primaryColor: defaultSettings.primaryColor,
  fixedHeader: defaultSettings.fixedHeader,
  fixSiderbar: defaultSettings.fixSiderbar,
  colorWeak: defaultSettings.colorWeak,

  hideHintAlert: false,
  hideCopyButton: false
})
const currentUser = reactive({
  nickname: 'Admin',
  avatar: 'A'
})
const breadcrumb = computed(() =>
  router.currentRoute.value.matched.concat().map((item) => {
    return {
      path: item.path,
      breadcrumbName: item.meta.title || ''
    }
  })
)

watchEffect(() => {
  if (router.currentRoute) {
    const matched = router.currentRoute.value.matched.concat()
    baseState.selectedKeys = matched.filter((r) => r.name !== 'index').map((r) => r.path)
    baseState.openKeys = matched
      .filter((r) => r.path !== router.currentRoute.value.path)
      .map((r) => r.path)
  }
})

onMounted(() => {
  setTimeout(() => {
    watermarkContent.value = 'New Mark'
  }, 2000)
})
</script>
