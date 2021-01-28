const path = require('path')
const defaultSettings = require('./src/settings.js')
const name = defaultSettings.title || 'vue Lego Admin' // page title

function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  //基础配置 详情看文档
  publicPath: './',
  outputDir: '../docker_console/vue_console',
  assetsDir: 'static',
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: false,
  devServer: {
    port: 8080,
    open: true,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VUE_APP_BASE_API]: { //需要代理的路径   例如 '/lego'
        target: `http://127.0.0.1:9567`, //代理到 目标路径
        changeOrigin: true,
        secure: false,
        logLevel: 'debug',
        pathRewrite: { // 修改路径数据
          ['^'+process.env.VUE_APP_BASE_API]: '' // 举例 '^/lego:""' 把路径中的/api字符串删除
        }
      }
    },
  },
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    devtool: 'source-map',
    name: name,
    resolve: {
      alias: {
        '@': resolve('src')
      }
    }
  },
}