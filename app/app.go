package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/TianQinS/crontab2/config"
)

func notFoundHandler(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"ok":   false,
		"data": "404",
	})
}

func forbiddenHandler(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"ok":   false,
		"data": "403",
	})
}

func NewApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnErrorCode(iris.StatusForbidden, forbiddenHandler)

	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableAutoFireStatusCode: false,
		Charset:                   config.Conf.Http.HttpCharset,
	}))
	InitApi(app)
	return app
}
