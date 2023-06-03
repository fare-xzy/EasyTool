import { resolve } from 'path'
import { externalizeDepsPlugin } from 'electron-vite'
import vue from '@vitejs/plugin-vue'

// 自动导入vue中hook reactive ref等
import AutoImport from 'unplugin-auto-import/vite'
//自动导入ui-组件 比如说ant-design-vue  element-plus等
//是此插件无法处理非组件模块，如 message，这种组件需要手动加载：
import Components from 'unplugin-vue-components/vite'
//ant-design-vue
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'

/** @type {import('vite').UserConfig} */
export default {
  main: {
    plugins: [externalizeDepsPlugin()]
  },
  preload: {
    plugins: [externalizeDepsPlugin()]
  },
  renderer: {
    server: {
      host: '0.0.0.0',
      port: 8991,
      // 是否开启 https
      https: false,
      // 设置反向代理，跨域
      proxy: {
        '/devApi': {
          // 后台地址
          target: 'http://127.0.0.1:9527/',
          changeOrigin: true,
          // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
          rewrite: (path: string) => path.replace(/^\/devApi/, '')
        },
        '/rap2': {
          target: 'http://rap2api.taobao.org/app/mock/310178/',
          changeOrigin: true,
          // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
          rewrite: (path: string) => path.replace(/^\/rap2/, '')
        }
      }
    },

    resolve: {
      alias: {
        '@renderer': resolve('src/renderer/src')
      }
    },
    plugins: [
      vue(),
      AutoImport({
        //安装两行后你会发现在组件中不用再导入ref，reactive等
        imports: ['vue', 'vue-router'],
        dts: 'src/auto-import.d.ts',
        //ant-design-vue
        resolvers: [AntDesignVueResolver()]
      }),
      Components({
        dirs: ['src/components'],
        //ant-design-vue   importStyle = false 样式就没了
        resolvers: [AntDesignVueResolver({ importStyle: true, resolveIcons: true })]
      })
    ],
    css: {
      preprocessorOptions: {
        less: {
          // DO NOT REMOVE THIS LINE
          javascriptEnabled: true,
          modifyVars: {
            // hack: `true; @import 'ant-design-vue/dist/antd.variable.less'`,
            // '@primary-color': '#eb2f96', // 全局主色
          }
        }
      }
    },
    optimizeDeps: {
      include: ['@ant-design/icons-vue', 'ant-design-vue']
    }
  }
}
