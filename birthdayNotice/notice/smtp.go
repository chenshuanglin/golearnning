package notice

import (
	"log"
	"net/smtp"
	"strings"
)

const (
	email  = "*****@163.com"
	passwd = "*****"
	host   = "smtp.163.com:25"
)

func SendMessage(user *User, message string) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", email, passwd, hp[0])
	to := strings.Split(user.Email, ";")
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"
	subject := "生日提醒"
	msg := []byte("To: " + user.Email + "\r\nFrom: " + email + "<" + email + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + message)
	err := smtp.SendMail(host, auth, email, to, msg)
	if err != nil {
		log.Printf("发送邮件失败，失败原因是:\n%v\n", err)
	}
}
