package Routes

import (
	"github.com/kataras/iris"
	"server/Controllers"
)

func Configure(app *iris.Application) {
	// test
	app.Post("/test", Controllers.Test)

	// Index
	app.Get("/", Controllers.GetIndexHandler)

	// Core
	app.Post("/Add", Controllers.Add)
	app.Post("/Reply", Controllers.Reply)
	app.Post("/Forget", Controllers.Forget)
	app.Post("/Status", Controllers.Status)
}
