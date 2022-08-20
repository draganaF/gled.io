package services

import (
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/draganaF/gled.io/EmailService/dto"
)

type EmailSender struct {
	smtpClient *mail.SMTPClient
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		smtpClient: CreateSMTPClient(),
	}
}

func (emailSender *EmailSender) Send(emailData *dto.Email) bool {

	email := mail.NewMSG()
	if len(emailData.From) != 0 {
		email.SetFrom(emailData.From)
	} else {
		email.SetFrom(os.Getenv("FROM"))
	}
	email.AddTo(emailData.To)
	email.SetSubject(emailData.Subject)
	email.SetBody(mail.TextPlain, emailData.Body)

	err := email.Send(emailSender.smtpClient)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
