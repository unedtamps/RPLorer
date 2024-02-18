package helper

import (
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/util"
	"gopkg.in/gomail.v2"
)

var Dialer *gomail.Dialer

type emailService struct {
	from    string
	to      string
	subject string
	body    string
}

func NewEmail(subject string, to string, body <-chan string) *emailService {
	return &emailService{
		from:    config.Env.EmailSender,
		to:      to,
		subject: subject,
		body:    <-body,
	}
}

func (e *emailService) Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", e.to)
	m.SetHeader("Subject", e.subject)
	m.SetBody("text/html", e.body)
	err := Dialer.DialAndSend(m)
	if err != nil {
		util.Log.Info("Error sending email: ", err)
	} else {
		util.Log.Info("Email sent successfully to ", e.to)
	}
}
