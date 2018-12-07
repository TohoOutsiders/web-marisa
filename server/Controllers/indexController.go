package Controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetIndexHandler(ctx iris.Context) {
	ctx.JSON(context.Map{
		"code": 200,
		"message": "hello Marisa~",
	})
}