/**
 * @Author: Tomonori
 * @Date: 2019/6/19 15:47
 * @File: routes
 * @Desc:
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/datasource"
	"server/repository"
	"server/service"
)

func Configure(app *gin.Engine) {
	// 实现服务接口
	db := datasource.GetInstace().GetMysqlDB()
	service := service.NewMemoriseService(
		repository.NewMemoriseRepo(db),
	)

	// 实例化控制器
	index := &controller.Index{}

	core := &controller.Core{Service: service}

	v1 := app.Group("/")
	{
		v1.GET("/", index.Get)

		v1.POST("/Add", core.Add)
		v1.POST("/Reply", core.Reply)
		v1.POST("/Forget", core.Forget)
		v1.POST("/Status", core.Status)
	}
}
