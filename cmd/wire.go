//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/app"
	"github.com/gutrse3321/aki/pkg/config"
	"github.com/gutrse3321/aki/pkg/log"
	"github.com/gutrse3321/aki/pkg/transports/http"
	"github.com/gutrse3321/web-marisa/src"
	"github.com/gutrse3321/web-marisa/src/controller"
)

/**
 * @Author: Tomonori
 * @Date: 2021/8/12
 * @Title:
 * --- --- ---
 * @Desc:
 */
var wireSet = wire.NewSet(
	log.WireSet,
	config.WireSet,
	http.WireSet,
	src.ProviderSet,
	controller.ProviderSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(wireSet))
}
