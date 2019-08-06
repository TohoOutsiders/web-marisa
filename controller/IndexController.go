/**
 * @Author: Tomonori
 * @Date: 2019/6/19 15:38
 * @File: IndexController
 * @Desc:
 */
package controller

import (
	"github.com/gin-gonic/gin"
)

type Index struct {
}

func (i *Index) Get(ctx *gin.Context) {
	Json(ctx, &ModelAndView{200, "Hello marisa~"})
}
