package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"log"
	"os"
	"web-marisa/server/Datasource"
	"web-marisa/server/Middlewares/setting"
	"web-marisa/server/Routes"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover2.New())
	app.Use(logger.New())
	app.Configure(Routes.Configure)

	issue := Datasource.GetInstace().InitDataPool();
	if !issue {
		log.Println("Inital database pool fail")
		os.Exit(1)
	}

	app.Run(
		iris.Addr(fmt.Sprintf(":%d", setting.HttpPort)),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
