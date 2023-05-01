package test

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Zhang HaoYu <310123665@qq.com>"
	e.To = []string{"310123665@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是<b>123456</b>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "310123665@qq.com", "vplxaeoiawtpbjda", "smtp.qq.com"))
	if err != nil {
		t.Fatal(err)
	}
}
