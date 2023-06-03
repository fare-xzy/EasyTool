<template>
  <div style="margin-right: 12px">
    <a-space>
      <a-dropdown>
        <template #overlay>
          <a-menu>
            <a-menu-item>
              <template #icon>
                <SettingOutlined />
              </template>
              <span>个人设置</span>
            </a-menu-item>
            <a-menu-item>
              <template #icon>
                <LogoutOutlined />
              </template>
              <span>退出登录</span>
            </a-menu-item>
          </a-menu>
        </template>
        <a-avatar shape="square" size="small">
          <template #icon>
            <UserOutlined />
          </template>
          {{ currentUser.nickname }}
        </a-avatar>
      </a-dropdown>
    </a-space>
  </div>
</template>

<script setup lang="ts">
export type CurrentUser = {
  nickname: string
  avatar?: string
}

defineProps<{
  currentUser: CurrentUser
}>()

</script>
