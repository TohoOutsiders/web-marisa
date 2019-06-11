package Controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"server/Models"
	"server/Services"
	"server/repository"
)

var (
	service = &Services.MemoriseService{
		Repo: &repository.MemoriseRepo{},
	}
)


func Add(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)

	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Println("Controller Add() error: %s", err)
	} else {
		if data := service.Add(memory); data != nil {
			ctx.JSON(context.Map{
				"code": 200,
				"data": data,
			})
		}
	}
}

func Reply(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Println("Controller Reply() error: %s", err)
	} else {
		code, data := service.Reply(memory)
		ctx.JSON(context.Map{
			"code": code,
			"data": data,
		})
	}
}

func Forget(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Println("Controller Forget() error: %s", err)
	} else {
		if flag := service.Forget(memory.Answer); flag {
			ctx.JSON(context.Map{
				"code": 200,
				"data": "success",
			})
		}
	}
}

func Status(ctx iris.Context) {
	count := service.Status()
	ctx.JSON(context.Map{
		"code": http.StatusOK,
		"data": count,
	})
}
