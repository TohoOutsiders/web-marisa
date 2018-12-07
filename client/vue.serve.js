const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = webpackConfig => {
  webpackConfig.plugin('html').tap(([options]) => [
    Object.assign(options, {
      minify: false,
      inject: true,
      chunksSortMode: 'none'
    })
  ])
}
