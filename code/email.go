package email

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/smtp"
)

//MailSender interface for sending emails
type MailSender interface {
	SendEmail(to string, subject string, body string)
}

type message struct {
	to   string
	body string
}

type mailSender struct {
	auth          smtp.Auth
	sendMail      func(addr string, a smtp.Auth, from string, to []string, msg []byte) error
	from          string
	serverAddress string
	enabled       bool
	ssl           bool
	channel       chan message
	logger        *logrus.Entry
}

// SendEmail sends given message via email.
func (ms mailSender) SendEmail(to string, subject string, body string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = fmt.Sprintf("Subject: %v\n", subject)
	ms.channel <- message{to, subject + mime + body}
	ms.logger.Info("added new message to chanel")
}

func (ms mailSender) send() {
	if ms.enabled {
		for {
			message := <-ms.channel
			ms.logger.Info("sending email to ", message.to)
			if err := ms.sendMail(
				ms.serverAddress,     // server address
				ms.auth,              // authentication
				ms.from,              // sender's address
				[]string{message.to}, // recipients' address
				[]byte(message.body), // message body
			); err != nil {
				ms.logger.Warn(errors.Wrapf(err, "failed to send email to %v", message.to))
			}
		}
	} else {
		for {
			message := <-channel
			ms.logger.Debug("sending emails disabled", message.to, message.body)
		}
	}

}

//NewMailSender constructor for async mail sender
func NewMailSender(log *logrus.Entry) MailSender {
	// the channel is buffered up to 100 messages
	var channel = make(chan message, 100)
	enabled := viper.GetBool("mail_enabled")
	port := viper.GetInt("mail_port")
	sender := mailSender{logger: log, enabled: enabled, channel: channel, ssl: port == 465}
	if enabled {
		sender.from = fmt.Sprintf("%v <%v>", viper.GetString("mail_from"), viper.GetString("mail_user"))
		host := viper.GetString("mail_server")
		sender.serverAddress = fmt.Sprintf("%v:%v", host, port)
		if sender.ssl {
			sender.sendMail = sendMailSSL
			sender.auth = loginAuth{viper.GetString("mail_user"), viper.GetString("mail_password"), host}
		} else {
			sender.sendMail = smtp.SendMail
			sender.auth = smtp.PlainAuth("", viper.GetString("mail_user"), viper.GetString("mail_password"), host)
		}
	}
	go sender.send()
	return &sender
}
