package smtp

import (
	"net/smtp"
	"os"
)

type SmtpSender interface {
	SendTestMail() error
}

type SmtpSenderBasic struct{}

func (s SmtpSenderBasic) SendTestMail() error {
	auth := smtp.PlainAuth("", os.Getenv(USER), os.Getenv(PASSWORD), os.Getenv(SERVER))

	msg := []byte("To: " + os.Getenv(USER) + "\r\n" +
		"Subject: Test mail for Guolmel\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	errSmtp := smtp.SendMail(os.Getenv(SERVER)+":587", auth, os.Getenv(USER),
		[]string{os.Getenv(USER)}, msg)

	if errSmtp != nil {
		return errSmtp
	}

	return nil
}
