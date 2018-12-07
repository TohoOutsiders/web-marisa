package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"web-marisa/server/Routes"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover2.New())
	app.Use(logger.New())
	app.Configure(Routes.Configure)

	app.Run(
		iris.Addr(":3000"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
