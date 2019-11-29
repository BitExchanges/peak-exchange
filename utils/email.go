package utils

import (
	"bytes"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"time"
)

// 邮件发送
// fromUser 发件人
// toUser   收件人
// subject  邮件主题
func SendEmail(fromUser, toUser, subject string) error {
	e := email.NewEmail()
	e.From = fromUser
	e.To = []string{toUser}
	e.Subject = subject
	t, err := template.ParseFiles("./email-template.html")
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
		IpAddress    string
	}{
		"农场主",
		"农民",
		time.Now().Format("2006/01/02"),
		"测试邮件",
		"192.168.0.1",
	})

	e.HTML = body.Bytes()
	//e.Attach(body,"email-template.html","text/html")
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", "769558579@qq.com", "fnbkjkhhmivzbebj", "smtp.qq.com"))
}
