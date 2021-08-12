package controller

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2021/8/12
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(
	NewIndexController,
	CreateInitControllersFn,
)
