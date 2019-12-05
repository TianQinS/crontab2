package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/TianQinS/crontab2/app"
	"github.com/TianQinS/crontab2/config"
	"github.com/TianQinS/crontab2/utils/mail"
	"github.com/TianQinS/fastapi/basic"
	"github.com/TianQinS/fastapi/timer"
	"github.com/kataras/iris"
)

var (
	port = flag.String("p", "23456", "iris http port")
)

// 加载crontab
func crontab() {
	for _, cmd := range config.Conf.Cmds {
		cronStr := cmd.CronStr
		title := cmd.Title
		execCmd := cmd.ExeCmd
		receivers := cmd.Receivers
		attach := cmd.AttachFile
		timer.AddCrontab(cronStr, title, func(args ...interface{}) {
			go func() {
				// 异步执行
				start := time.Now()
				log.Printf("%+v\n", cmd)
				if cmd.Valid {
					if res, err := basic.Exec(execCmd); err != nil {
						mail.SendMsg(title, fmt.Sprintf("指令执行失败: %+v。\n", err), execCmd, attach, receivers)
					} else {
						execCmd = fmt.Sprintf("执行时长: %d 执行指令: %s", time.Now().Sub(start).Seconds(), execCmd)
						mail.SendMsg(title, string(res), execCmd, attach, receivers)
					}
				} else {
					mail.SendMsg(title, "该项配置未开启。\n", execCmd, "", receivers)
				}
			}()
		})
	}
}

// 加载另外的httpHandle
func InitApi(app *iris.Application) {
}

func init() {
	crontab()
}

func main() {
	// runtime.GOMAXPROCS(2)
	flag.Parse()
	http := app.NewApp()
	InitApi(http)
	http.Run(iris.Addr(":" + *port))
}
