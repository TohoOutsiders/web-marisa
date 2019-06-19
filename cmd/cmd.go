/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:37
 * @File: cmd
 * @Desc:
 */
package cmd

import (
	"github.com/gin-gonic/gin"
	"server/datasource"
	"server/routes"
)

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	// 连接数据库
	datasource.ConnectDatabase()
	// 注册路由
	routes.Configure(r)

	return r
}
