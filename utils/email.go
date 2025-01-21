package utils

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendMessage(subject, content, to string) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <187046@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{to}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1870489@qq.com", "oaujtqx", "smtp.qq.com"))
	if err != nil {
		log.Println(err)
	}
}
