const vueBuild = require('./vue.build')
const vueServe = require('./vue.serve')

module.exports = {
  lintOnSave: true,
  css: {
    extract: false, // 不提取css
    sourceMap: process.env.NODE_ENV === 'production' ? false : true,
    loaderOptions: {
      stylus: {
        javascriptEnabled: true
      }
    }
  },
  devServer: {
    open: true,
    host: '0.0.0.0',
    port: 8888, // 可选
    https: false,
    hotOnly: false,
    proxy: null
  },
  productionSourceMap: process.env.NODE_ENV === 'production' ? false : true,
  chainWebpack: webpackConfig => {
    // webpackConfig.externals({
    //   vue: 'vue',
    //   'vue-router': 'VueRouter',
    //   'vue-ls': 'VueStorage',
    //   vuex: 'Vuex',
    //   axios: 'axios'
    // })
    webpackConfig.module
      .rule('images')
      .test(/\.(png|jpe?g|gif|webp)(\?.*)?$/)
      .use('url-loader')
      .loader('url-loader')
      .options({
        limit: 10000,
        name: 'image/[name].[hash:8].[ext]' // 这里将默认的img文件夹名改为image
      })
    webpackConfig.module
      .rule('svg')
      .test(/\.(svg)(\?.*)?$/)
      .use('file-loader')
      .loader('file-loader')
      .options({
        name: 'image/[name].[hash:8].[ext]' // 这里将默认的img文件夹名改为image
      })
    process.env.NODE_ENV === 'production' ? vueBuild(webpackConfig) : vueServe(webpackConfig)
  }
}
