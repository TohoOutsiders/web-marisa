package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"log"
	"os"
	"server/Datasource"
	"server/Middlewares/setting"
	"server/Routes"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("info")
	app.Use(recover2.New())
	app.Use(logger.New())

	issue := Datasource.GetInstace().InitDataPool()
	if !issue {
		log.Println("Inital database pool fail")
		os.Exit(1)
	}
	// 存在hero依赖注入，需要数据库连接成功，才开始注入
	app.Configure(Routes.Configure)

	app.Run(
		iris.Addr(fmt.Sprintf(":%d", setting.HttpPort)),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
