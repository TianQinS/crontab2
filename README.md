# Crontab2
[![GoDoc](https://godoc.org/github.com/TianQinS/crontab22?status.svg)](https://godoc.org/github.com/TianQinS/crontab22)

**特点**

类似unix crontab的计划任务功能，支持屏显信息和运行日志的邮件发送。
>1. 配置样式：
	{
		Title = "测试指令",
		ExeCmd = "echo 'test'",
		CronStr = "*/2 * * * *",
		Valid = false,
		Receivers = [],
		AttachFile = ""
	}

