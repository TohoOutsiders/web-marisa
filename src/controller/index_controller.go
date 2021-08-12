package controller

import (
	"github.com/gin-gonic/gin"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2021/8/12
 * @Title:
 * --- --- ---
 * @Desc:
 */

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func CreateInitControllersFn(index *IndexController) httpServer.InitControllers {
	return func(g *gin.RouterGroup) {
		g.GET("/", index.Home)
	}
}

func (i *IndexController) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &gin.H{
		"message": "お久しぶりです",
	})
}
