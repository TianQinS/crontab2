package mail

import (
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/TianQinS/crontab2/config"
	"gopkg.in/gomail.v2"
)

var (
	Conf = &config.Conf.Mail
)

/* 发送crontab指令执行信息邮件
@param title: 邮件标题主体
@param msg: 指令执行屏显信息，尽量不要太多
@param cmd: 执行的指令信息
*/
func SendMsg(title, msg, cmd, attach string, receivers []string) {
	num := strings.Count(msg, "\n")
	if num < 2 {
		num = 2
	}

	height := strconv.Itoa(num * MSG_LINE_HEIGHT)
	content := fmt.Sprintf(MSG_INFO, MSG_CSS, height, msg, cmd)
	receivers = append(receivers, Conf.DefaultReceivers...)
	SendMail(content, receivers, title, attach)
}

// Send normal mail.
func SendMail(content string, receivers []string, title, attach string) {
	if len(receivers) == 0 {
		// 如果没有配置接收邮件者
		return
	}
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", Conf.User)
		m.SetHeader("To", receivers...)
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", title)
		m.SetBody("text/html", content)
		if len(attach) > 0 {
			if _, err := os.Stat(attach); err == nil {
				m.Attach(attach)
			}
		}

		d := gomail.NewDialer(Conf.Host, Conf.Port, Conf.User, Conf.Pwd)
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err.Error())
			debug.PrintStack()
		}
	}()
}
