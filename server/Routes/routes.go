package Routes

import (
	"github.com/kataras/iris"
	"web-marisa/server/Controllers"
)

func Configure(app *iris.Application) {
	// Index
	app.Get("/", Controllers.GetIndexHandler)

	// Core
	app.Post("/Add", Controllers.Add)
}
