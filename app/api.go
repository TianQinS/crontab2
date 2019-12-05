package app

import (
	"net/http"
	"net/http/pprof"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	"github.com/TianQinS/crontab2/config"
)

// The relevant API functions of cache.
func InitApi(app *iris.Application) {
	// var getAPI router.Party

	if config.Conf.Debug {
		ppApi := app.Party("/debug")
		ppApi.Get("/pprof", pprofHandler(pprof.Index))
		ppApi.Get("/cmdline", pprofHandler(pprof.Cmdline))
		ppApi.Get("/profile", pprofHandler(pprof.Profile))
		ppApi.Post("/symbol", pprofHandler(pprof.Symbol))
		ppApi.Get("/symbol", pprofHandler(pprof.Symbol))
		ppApi.Get("/trace", pprofHandler(pprof.Trace))
		ppApi.Get("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		ppApi.Get("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		ppApi.Get("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		ppApi.Get("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		ppApi.Get("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		ppApi.Get("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))

		getAPI := app.Party("/get")
		getAPI.Get("/crontab", ShowCrontab)
	}
}

func pprofHandler(f http.HandlerFunc) context.Handler {
	handler := http.HandlerFunc(f)
	return func(ctx iris.Context) {
		handler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	}
}
