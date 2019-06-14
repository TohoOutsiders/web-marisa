package Controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"net/http"
	"server/Models"
	"server/Services"
)

func Add(ctx iris.Context, service Services.IMemoriseService) ModelAndView {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)

	if err != nil {
		fmt.Println("Controller Add() error: %s", err)
		return ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		}
	}
	if data := service.Add(memory); data != nil {
		return ModelAndView{
			Code: http.StatusOK,
			Data: data,
		}
	}
	return ModelAndView{
		Code: http.StatusBadGateway,
		Data: "服务器繁忙",
	}
}

func Reply(ctx iris.Context, service Services.IMemoriseService) ModelAndView {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		fmt.Println("Controller Reply() error: %s", err)
		return ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		}
	}
	code, data := service.Reply(memory)
	return ModelAndView{
		Code: code,
		Data: data,
	}
}

func Forget(ctx iris.Context, service Services.IMemoriseService) ModelAndView {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		fmt.Println("Controller Forget() error: %s", err)
		return ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		}
	}
	if flag := service.Forget(memory.Answer); flag {
		return ModelAndView{
			Code: http.StatusOK,
			Data: "success",
		}
	}
	return ModelAndView{
		Code: http.StatusBadGateway,
		Data: "服务器繁忙",
	}
}

func Status(service Services.IMemoriseService) ModelAndView {
	count := service.Status()
	return ModelAndView{
		Code: http.StatusOK,
		Data: count,
	}
}
