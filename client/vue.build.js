const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = webpackConfig => {
  webpackConfig.baseUrl = './static' // CDN路径
  webpackConfig.plugin('html').tap(([options]) => [
    Object.assign(options, {
      minify: {
        removeComments: true, // 是否去掉注释
        removeCommentsFromCDATA: true, // 移除来自字符数据的注释
        collapseWhitespace: true, // 是否去掉空格
        conservativeCollapse: false, // 删除空格，总是保留一个空格
        collapseInlineTagWhitespace: true, // 去除内联标签中的空格
        collapseBooleanAttributes: true, // 简化布尔属性
        removeRedundantAttributes: true, // 删除多余的属性
        removeAttributeQuotes: false, // 是否移除属性引号
        removeEmptyAttributes: true, // 移除空属性
        removeScriptTypeAttributes: true,
        removeStyleLinkTypeAttributes: true,
        useShortDoctype: true,
        minifyJS: true, // 是否压缩html里的js
        minifyCSS: true // 是否压缩html里的css
      },
      inject: true,
      chunksSortMode: 'none'
    })
  ])
}
