package Controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"web-marisa/server/Models"
	"web-marisa/server/Services"
)

func Add(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)

	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Errorf("Controller Add() error: %s", err)
	} else {
		data := make(map[string]interface{})
		data["ip"] = memory.Ip
		data["keyword"] = memory.Keyword
		data["answer"] = memory.Answer
		if Services.AddMemory(data) {
			ctx.JSON(context.Map{
				"code": 200,
				"data": data,
			})
		}
	}
}
