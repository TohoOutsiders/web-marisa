/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:37
 * @File: cmd
 * @Desc:
 */
package cmd

import (
	"github.com/gin-gonic/gin"
	"server/common/segment"
	"server/routes"
)

var err error

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	//加载词典
	segment.Init()

	// 连接数据库 & 依赖注入 & 注册路由
	routes.Configure(r)

	return r
}
