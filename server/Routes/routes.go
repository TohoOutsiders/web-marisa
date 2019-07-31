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
	//app.UseGlobal(before)
	app.PartyFunc("/", func(r iris.Party) {
		r.Post("Add", hero.Handler(Controllers.Add))
		r.Post("Reply", hero.Handler(Controllers.Reply))
		r.Post("Forget", hero.Handler(Controllers.Forget))
		r.Post("Status", hero.Handler(Controllers.Status))
	})
}

//func before(ctx iris.Context) {
//	log.Println("fuck: ", ctx.Path())
//}
