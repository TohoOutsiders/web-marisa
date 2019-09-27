/**
 * @Author: Tomonori
 * @Date: 2019/6/19 15:59
 * @File: CoreController
 * @Desc:
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/models"
	"server/service"
)

type Core struct {
	Service service.IMemoriseService `inject:""`
}

func (c *Core) Add(ctx *gin.Context) {
	var memory models.Memorise

	err := ctx.ShouldBind(&memory)

	if err != nil {
		log.Println("[Controller] Add() error: ", err)
		Json(ctx, &ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		})
	}
	if data := c.Service.Add(memory); data != nil {
		Json(ctx, &ModelAndView{
			Code: http.StatusOK,
			Data: data,
		})
	}
}

func (c *Core) Reply(ctx *gin.Context) {
	var memory models.Memorise

	err := ctx.ShouldBind(&memory)
	if err != nil {
		log.Println("[Controller] Reply() error: ", err)
		Json(ctx, &ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		})
	}
	code, data := c.Service.Reply(memory)
	Json(ctx, &ModelAndView{
		Code: code,
		Data: data,
	})
}

func (c *Core) Forget(ctx *gin.Context) {
	var memory models.Memorise

	err := ctx.ShouldBind(&memory)
	if err != nil {
		log.Println("[Controller] Forget() error: ", err)
		Json(ctx, &ModelAndView{
			Code: http.StatusBadRequest,
			Data: err.Error(),
		})
	}
	if flag := c.Service.Forget(memory.Answer); flag {
		Json(ctx, &ModelAndView{
			Code: http.StatusOK,
			Data: "success",
		})
	} else {
		Json(ctx, &ModelAndView{
			Code: http.StatusBadGateway,
			Data: "服务器繁忙",
		})
	}
}

func (c *Core) Status(ctx *gin.Context) {
	count := c.Service.Status()
	Json(ctx, &ModelAndView{
		Code: http.StatusOK,
		Data: count,
	})
}
