package Routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"server/Controllers"
	"server/Datasource"
	"server/Services"
	"server/repository"
)

func Configure(app *iris.Application) {
	// hero
	db := Datasource.GetInstace().GetMysqlDB()
	hero.Register(
		Services.NewMemoriseService(
			repository.NewMemoriseRepo(db),
		),
	)

	// Index
	app.Get("/", Controllers.GetIndexHandler)

	// Core
	app.PartyFunc("/", func(r iris.Party) {
		app.Post("/Add", hero.Handler(Controllers.Add))
		app.Post("/Reply", hero.Handler(Controllers.Reply))
		app.Post("/Forget", hero.Handler(Controllers.Forget))
		app.Post("/Status", hero.Handler(Controllers.Status))
	})
}
