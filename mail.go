package robi

import (
	"crypto/tls"
	"fmt"
	"path/filepath"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mailer struct {
	Client *mail.SMTPServer
}

func (mailer *Mailer) String() string {

	return fmt.Sprintf("(%v,%v,%v)", mailer.Client.Host, mailer.Client.Port, mailer.Client.Username)

}

func NewMailer(host string, port int, user, password string) *Mailer {
	client := mail.NewSMTPClient()
	client.Host = host
	client.Port = port
	client.Username = user
	client.Password = password
	client.Encryption = mail.EncryptionSTARTTLS
	client.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mailer{Client: client}
}

func (mailer *Mailer) Send(from string, to []string, cc []string, subject string, body string, attaches []string) error {

	email := mail.NewMSG()

	if from == "" {
		email.SetFrom(mailer.Client.Username)
	} else {
		email.SetFrom(from)
	}

	for _, t := range to {
		email.AddTo(t)
	}
	for _, c := range cc {
		email.AddCc(c)
	}

	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	for _, attach := range attaches {
		email.Attach(&mail.File{FilePath: attach, Name: filepath.Base(attach), Inline: false})
	}

	if email.Error != nil {
		return email.Error
	}

	smtpClient, err := mailer.Client.Connect()

	if err != nil {
		return err
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	} else {
		return nil
	}

}
