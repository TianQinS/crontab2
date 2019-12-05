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

// 计划任务执行
func task(cmd config.Cmd) {
	title := cmd.Title
	execCmd := cmd.ExeCmd
	receivers := cmd.Receivers
	attach := cmd.AttachFile
	start := time.Now()
	log.Printf("%+v\n", cmd)
	if cmd.Valid {
		if res, err := basic.Exec(execCmd); err != nil {
			mail.SendMsg(title, fmt.Sprintf("指令执行失败: %+v %+v。\n", res, err), execCmd, attach, receivers)
		} else {
			execCmd = fmt.Sprintf("执行时长: %vs 执行指令: %s", time.Now().Sub(start).Seconds(), execCmd)
			mail.SendMsg(title, string(res), execCmd, attach, receivers)
		}
	} else {
		mail.SendMsg(title, "该项配置未开启。\n", execCmd, "", receivers)
	}
}

// 计划任务配置
func crontab(cmd config.Cmd) {
	cronStr := cmd.CronStr
	title := cmd.Title
	timer.AddCrontab(cronStr, title, func(args ...interface{}) {
		// 异步执行
		go task(cmd)
	})
}

// 加载另外的httpHandle
func InitApi(app *iris.Application) {
}

// 加载crontab
func init() {
	for _, cmd := range config.Conf.Cmds {
		crontab(cmd)
	}
}

func main() {
	// runtime.GOMAXPROCS(2)
	flag.Parse()
	http := app.NewApp()
	InitApi(http)
	http.Run(iris.Addr(":" + *port))
}
