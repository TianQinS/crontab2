package mail

import (
	"log"
	"testing"
	"time"
	//"github.com/stretchr/testify/assert"
)

func TestMsg(t *testing.T) {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Fatalln(e.Error())
		}
	}()
	SendMsg("测试指令", "测试指令执行内容\n行1\n行2", "测试指令内容", "", []string{"zdn3039@corp.netease.com"})
	time.Sleep(2 * time.Second)
}
