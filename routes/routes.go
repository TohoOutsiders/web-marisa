/**
 * @Author: Tomonori
 * @Date: 2019/6/19 15:47
 * @File: routes
 * @Desc:
 */
package routes

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"log"
	"server/controller"
	"server/datasource"
	"server/repository"
	"server/service"
)

func Configure(app *gin.Engine) {
	var err error

	// 实例化控制器
	index := &controller.Index{}
	// 控制器注入声明
	var core controller.Core

	// 连接数据库
	db := datasource.Db{}
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal: ", err)
	}

	// 依赖注入
	var ninject inject.Graph
	err = ninject.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &repository.MemoriseRepo{}},
		&inject.Object{Value: &service.MemoriseService{}},
		&inject.Object{Value: &core},
	)

	if err != nil {
		log.Fatal("inject fatal: ", err)
	}

	if err := ninject.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	v1 := app.Group("/")
	{
		v1.GET("/", index.Get)

		v1.POST("/Add", core.Add)
		v1.POST("/Reply", core.Reply)
		v1.POST("/Forget", core.Forget)
		v1.POST("/Status", core.Status)
	}
}
