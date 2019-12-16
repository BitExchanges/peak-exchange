package utils

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"time"
)

// 邮件发送
// fromUser 发件人
// toUser   收件人
// subject  邮件主题
func SendEmail(fromUser, toUser, subject, ip string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("TopEx-Email-Service <%s>", fromUser)
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
		time.Now().Format("2006/01/02 15:04:05"),
		"测试邮件",
		ip,
	})

	e.HTML = body.Bytes()
	//e.Attach(body,"email-template.html","text/html")
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", "769558579@qq.com", "fnbkjkhhmivzbebj", "smtp.qq.com"))
}

// 发送邮箱验证码
func SendCaptchaEmail(fromUser, toUser, captcha string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("TopEx-Email-Service <%s>", fromUser)
	e.To = []string{toUser}
	e.Subject = "验证码"
	t, err := template.ParseFiles("./captcha-template.html")
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
	}{
		fromUser,
		toUser,
		time.Now().Format("2006/01/02 15:04:05"),
		captcha,
	})

	e.HTML = body.Bytes()
	//添加附件
	//e.Attach(body,"email-template.html","text/html")
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", "769558579@qq.com", "fnbkjkhhmivzbebj", "smtp.qq.com"))

}

func SendEmails(toUser, subject, message, title, head, content, ip string) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("TopEx-Email <%s>", "769558579@qq.com")
	e.To = []string{toUser}
	e.Subject = subject
	t, err := template.ParseFiles("./common-template.html")
	if err != nil {
		log.Println("邮件模板加载失败: ", err)
		return
	}
	body := new(bytes.Buffer)
	t.Execute(body, struct {
		FromUser string
		ToUser   string
		TimeDate string
		Message  string
		Ip       string
		Title    string
		Head     string
		Content  string
	}{
		"769558579@qq.com",
		toUser,
		time.Now().Format("2006/01/02 15:04:05"),
		message,
		ip,
		title,
		head,
		content,
	})
	e.HTML = body.Bytes()
	err = e.Send("smtp.qq.com:587", smtp.PlainAuth("", "769558579@qq.com", "fnbkjkhhmivzbebj", "smtp.qq.com"))
	if err != nil {
		log.Println("邮件发送失败: ", err)
	}
}
