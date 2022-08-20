package services

import (
	"os"
	"strconv"

	mail "github.com/xhit/go-simple-mail/v2"
)

func CreateSMTPClient() *mail.SMTPClient {

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	server := mail.NewSMTPClient()
	server.Host = os.Getenv("HOST")
	server.Port = port
	server.Username = os.Getenv("API_KEY")
	server.Password = os.Getenv("SECRET")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		panic(err)
	}

	return smtpClient
}
