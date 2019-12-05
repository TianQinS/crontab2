package app

import (
	"github.com/kataras/iris"
	"github.com/TianQinS/crontab2/config"
)

// 查看当前Crontab配置。
func ShowCrontab(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"ok":   true,
		"data": config.Conf.Cmds,
	})
}
