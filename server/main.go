package main

import (
	"fmt"
	"github.com/kataras/iris"
	"server/Datasource"
	"server/Middlewares/setting"
	"server/Routes"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("info")

	// 连接数据库
	Datasource.ConnectDatabase(app)
	// 存在hero依赖注入，需要数据库连接成功，才开始注入
	app.Configure(Routes.Configure)

	app.Run(
		iris.Addr(fmt.Sprintf(":%d", setting.HttpPort)),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
