package help

import (
	"github.com/astaxie/beego"
	"net/smtp"
	"strings"
)

func SendMail(to, subject, body, mtype string) {
	user := beego.AppConfig.String("smtp.user")
	password := beego.AppConfig.String("smtp.password")
	host := beego.AppConfig.String("smtp.host")
	port := beego.AppConfig.String("smtp.port")

	auth := smtp.PlainAuth("", user, password, host)

	from := user
	sendTo := strings.Split(to, ",")
	var content_type string
	if mtype == "html" {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}
	body = "To: " + to + "\r\nFrom: " + from + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body
	msg := []byte(body)

	err := smtp.SendMail(host+":"+port, auth, from, sendTo, msg)
	Error(err)
}
