package main

import (
	"flag"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/TianQinS/crontab2/app"
	"github.com/TianQinS/crontab2/config"
	"github.com/TianQinS/crontab2/utils/mail"
	"github.com/TianQinS/fastapi/basic"
	"github.com/TianQinS/fastapi/timer"
	"github.com/axgle/mahonia"
	"github.com/kataras/iris"
)

var (
	port = flag.String("p", "23456", "iris http port")
	// 当前执行任务数
	gCnt int64 = 0
)

// 针对gbk编码的系统
func byte2string(out []byte) string {
	data := string(out)
	if config.Conf.CharSet == "gbk" {
		srcCoder := mahonia.NewDecoder("gbk")
		srcResult := srcCoder.ConvertString(data)
		tagCoder := mahonia.NewDecoder("utf8")
		_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
		data = string(cdata)
	}
	return data
}

// 计划任务执行
func task(cmd config.Cmd) {
	title := cmd.Title
	execCmd := cmd.ExeCmd
	receivers := cmd.Receivers
	attach := cmd.AttachFile
	start := time.Now()
	atomic.AddInt64(&gCnt, 1)
	log.Printf("[Cnt] %d start %+v\n", gCnt, cmd)
	if cmd.Valid {
		if res, err := basic.Exec(execCmd); err != nil {
			// 错误信息添加默认收件人
			receivers = append(receivers, config.Conf.Mail.DefaultReceivers...)
			mail.SendMsg(title, fmt.Sprintf("[CMD] Failed: %+v %s. \n", err, byte2string(res)), execCmd, attach, receivers)
		} else {
			execCmd = fmt.Sprintf("[TimeUsed]: %vs [CMD]: %s", time.Now().Sub(start).Seconds(), execCmd)
			mail.SendMsg(title, byte2string(res), execCmd, attach, receivers)
		}
	} else {
		mail.SendMsg(title, "[CMD] Invalid.\n", execCmd, "", receivers)
	}
	atomic.AddInt64(&gCnt, -1)
	log.Printf("[Cnt] %d end %s\n", gCnt, title)
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
	http.Run(iris.Addr(":" + *port))
}
