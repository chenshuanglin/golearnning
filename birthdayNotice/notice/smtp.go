package notice

import (
	"log"
	"net/smtp"
	"strings"
)

const (
	email  = "********@163.com"
	passwd = "*******"
	host   = "smtp.163.com:25"
)

type mes struct {
	to          string
	from        string
	fromEmail   string
	subject     string
	contentType string
	body        string
}

func (m *mes) String() string {
	return "To: " + m.to + "\r\nFrom: " + m.from + "<" + m.fromEmail + ">\r\nSubject: " + m.subject + "\r\n" + m.contentType + "\r\n\r\n" + m.body
}

func sendMessage(user *User, message string) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", email, passwd, hp[0])
	to := strings.Split(user.Email, ";")
	contentType := "Content-Type: text/plain; charset=UTF-8"
	subject := "生日提醒"
	m := &mes{
		to:          user.Email,
		from:        "rake的生日提醒",
		fromEmail:   email,
		subject:     subject,
		contentType: contentType,
		body:        message,
	}
	msg := []byte(m.String())
	err := smtp.SendMail(host, auth, email, to, msg)
	if err != nil {
		log.Printf("发送邮件失败，失败原因是:\n%v\n", err)
	}
}
