/**
 * @Author: Tomonori
 * @Date: 2019/6/19 16:12
 * @File: Constant
 * @Desc:
 */
package controller

import "github.com/gin-gonic/gin"

type ModelAndView struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Json(ctx *gin.Context, model *ModelAndView) {
	ctx.JSON(model.Code, &model)
}
