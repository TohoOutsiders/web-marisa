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
	"server/common/cache"
	"server/common/datasource"
	"server/common/rabbitmq"
	"server/controller"
	"server/repository"
	"server/service"
)

func Configure(app *gin.Engine) {
	var err error

	// 实例化控制器
	index := &controller.Index{}
	// 控制器注入声明
	var core controller.Core

	// 注入声明
	db := datasource.Db{}
	redis := cache.Redis{}
	rabbit := rabbitmq.Mq{}

	// 注入
	var injector inject.Graph
	err = injector.Provide(
		&inject.Object{Value: &core},
		&inject.Object{Value: &db},
		&inject.Object{Value: &redis},
		&inject.Object{Value: &rabbit},
		&inject.Object{Value: &repository.MemoriseRepo{}},
		&inject.Object{Value: &service.MemoriseService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//数据库连接
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal:", err)
	}
	//连接缓存服务器
	err = redis.Connect()
	if err != nil {
		log.Fatal("redis fatal:", err)
	}
	// 连接rabbitmq
	err = rabbit.Connect()
	if err != nil {
		log.Fatal("RabbitMQ fatal:", err)
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
