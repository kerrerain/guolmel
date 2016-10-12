package smtp

import (
	"github.com/magleff/guolmel/models"
	"net/smtp"
	"os"
)

type SmtpSender interface {
	SendTestMail() error
}

type SmtpSenderBasic struct{}

func (s SmtpSenderBasic) SendMail(msg []byte) error {
	auth := smtp.PlainAuth("", os.Getenv(USER), os.Getenv(PASSWORD), os.Getenv(SERVER))

	errSmtp := smtp.SendMail(os.Getenv(SERVER)+":587", auth, os.Getenv(USER),
		[]string{os.Getenv(USER)}, msg)

	if errSmtp != nil {
		return errSmtp
	}

	return nil
}

func (s SmtpSenderBasic) SendTestMail() error {
	msg := []byte("To: " + os.Getenv(USER) + "\r\n" +
		"Subject: Test mail for Guolmel\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	return s.SendMail(msg)
}

func (s SmtpSenderBasic) SendBudgetState(budget models.Budget) error {
	msg := []byte("To: " + os.Getenv(USER) + "\r\n" +
		"Subject: Budget state at " + budget.LastModificationDate.String() + "\r\n" +
		"\r\n" +
		budget.ToMessage())

	return s.SendMail(msg)
}
