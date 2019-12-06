package config

import (
	"log"
	"runtime/debug"

	"github.com/lytics/confl"
)

const (
	// json格式配置文件
	CONF_FILE = "./confSample.ini"
)

var (
	Conf = new(Config)
)

type Config struct {
	Debug   bool   // 是否调试模式
	CharSet string // 操作系统编码
	Mail    Mail
	Http    Http
	Cmds    []Cmd
}

// 服务邮件配置
type Mail struct {
	Host             string // smtp.163.com
	Port             int    // 25
	User             string
	Pwd              string
	DefaultReceivers []string // 邮件默认接收者列表
}

// Http相关
type Http struct {
	HttpCharset string
}

// 运行指令相关
type Cmd struct {
	Title      string   // title key
	ExeCmd     string   // 执行指令
	CronStr    string   // crontab字符串
	Valid      bool     // 是否生效
	Receivers  []string // 邮件接受者
	AttachFile string   // 可能需要发送的附件
}

func init() {
	if _, err := confl.DecodeFile(CONF_FILE, Conf); err != nil {
		log.Printf("[Config] read conf file error. info=%s trace=%s\n", err.Error(), string(debug.Stack()))
	}
}
