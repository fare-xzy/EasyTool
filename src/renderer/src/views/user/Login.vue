<template>
  <div class="main">
    <a-form id="formLogin" :model="userInfo" class="user-layout-login" @submit="handleSubmit">
      <a-tabs
        :active-key="customActiveKey"
        :tab-bar-style="{ textAlign: 'center', borderBottom: 'unset' }"
        :centered="true"
      >
        <a-tab-pane key="tab1" tab="账户密码登录">
          <a-form-item
            name="Host"
            :rules="[
              {
                required: true,
                pattern:
                  /^(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)$/,
                message: '请输入正确的服务器IP地址'
              }
            ]"
          >
            <a-input v-model:value="userInfo.Host" size="large" type="text" placeholder="服务器IP">
              <template #prefix>
                <hdd-outlined :style="{ color: 'rgba(0,0,0,.25)' }" />
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
            name="Port"
            :rules="[
              {
                required: true,
                pattern: /^[0-9]*$/,
                message: '请输入正确的SSH端口'
              }
            ]"
          >
            <a-input v-model:value="userInfo.Port" size="large" type="text" placeholder="端口">
              <template #prefix>
                <link-outlined :style="{ color: 'rgba(0,0,0,.25)' }" />
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
            name="UserName"
            :rules="[
              {
                required: true,
                message: '请输入用户名'
              }
            ]"
          >
            <a-input
              v-model:value="userInfo.UserName"
              size="large"
              type="text"
              placeholder="用户名"
            >
              <codepen-outlined />
              <template #prefix>
                <user-outlined :style="{ color: 'rgba(0,0,0,.25)' }" />
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
            name="Port"
            :rules="[
              {
                required: true,
                message: '请输入密码'
              }
            ]"
          >
            <a-input-password v-model:value="userInfo.Password" size="large" placeholder="密码">
              <template #prefix>
                <lock-outlined :style="{ color: 'rgba(0,0,0,.25)' }" />
              </template>
            </a-input-password>
          </a-form-item>
        </a-tab-pane>
        <!--        <a-tab-pane key="tab2" :tab="$t('user.login.tab-login-quick')">-->
        <!--          <a-form-item>-->
        <!--            <a-input disabled="true" size="large" type="text" placeholder="敬请期待" v-decorator="['undone']">-->
        <!--              <a-icon slot="prefix" type="mobile" :style="{ color: 'rgba(0,0,0,.25)' }"/>-->
        <!--            </a-input>-->
        <!--          </a-form-item>-->
        <!--        </a-tab-pane>-->
      </a-tabs>

      <a-form-item style="margin-top: 24px">
        <a-button
          size="large"
          type="primary"
          html-type="submit"
          class="login-button"
          :loading="status.loginLoading"
          >登录</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script lang="ts" setup>
import { UserAuthStore } from '@renderer/store/user'
const { proxy } = getCurrentInstance()
const customActiveKey = 'tab1'
const status = reactive({
  loginLoading: false
})
const userInfo = reactive({
  Host: '',
  Port: '',
  UserName: '',
  Password: ''
})

const user = UserAuthStore()

function handleSubmit(): void {
  status.loginLoading = true
  user
    .Login(userInfo)
    .then((res) => {
      if (res.code == 200) {
        loginSuccess()
      } else {
        loginFailed(res.message)
      }
    })
    .finally(() => {
      status.loginLoading = false
    })
}

/**
 * 登录成功函数
 */
function loginSuccess(): void {
  // const n = Math.floor(Math.random() * 16.0).toString(16)
  // storage.set(ACCESS_TOKEN, n, new Date().getTime() + 7 * 24 * 60 * 60 * 1000)
  user.isAuthenticated = true
  proxy.$router.push({ path: '/list/table-list' })
  // 延迟 1 秒显示欢迎信息
  setTimeout(() => {
    proxy.$notification['success']({
      message: '欢迎',
      description: `${proxy.$timeFix()}，欢迎回来`
    })
  }, 1000)
}

function loginFailed(data: string): void {
  // this.isLoginError = true
  proxy.$notification['error']({
    message: '错误',
    description: data,
    duration: 4
  })
}
</script>

<style lang="less" scoped>
.user-layout-login {
  label {
    font-size: 14px;
  }

  .forge-password {
    font-size: 14px;
  }

  button.login-button {
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}
</style>
