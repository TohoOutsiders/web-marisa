package Routes

import (
	"github.com/kataras/iris"
	"web-marisa/server/Controllers"
)

func Configure(app *iris.Application) {
	app.Get("/", Controllers.GetIndexHandler)
}
