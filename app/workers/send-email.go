package workers

import (
	"github.com/Focinfi/sakura/config"
	"github.com/Focinfi/sakura/libs/worker"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/gomail.v2"
)

// Emailer sends emails
type Emailer struct {
	smtp     string
	from     string
	password string
}

var emailer = Emailer{
	smtp:     config.Config.EmailSMTP,
	from:     config.Config.EmailUser,
	password: config.Config.EmailPassword,
}

// SendEmail sends email
func (emailer *Emailer) SendEmail(subject, body string, to ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailer.from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(emailer.smtp, 587, emailer.from, emailer.password)
	return d.DialAndSend(m)
}

// SendEmail sends email in background
func SendEmail(subject, body string, to ...string) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		err := emailer.SendEmail(subject, body, to...)
		if wu.IsCancelled() {
			return nil, nil
		}

		return nil, err
	}
}

// NewEmailWorker worker
func NewEmailWorker() *worker.Worker {
	return worker.
		New(defaultQueue).
		SetRetry(true)
}

// EmailWorker for send email worker
var EmailWorker = NewEmailWorker()
