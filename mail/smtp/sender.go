package smtp

import (
	"github.com/magleff/guolmel/models"
	"net/smtp"
	"os"
	"time"
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

func (s SmtpSenderBasic) SendBudgetMessage(budget models.Budget, subject string) error {
	msg := []byte("To: " + os.Getenv(USER) + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		budget.ToMessage())

	return s.SendMail(msg)
}

func (s SmtpSenderBasic) SendBudgetState(budget models.Budget) error {
	return s.SendBudgetMessage(budget, models.BUDGET_STATE_FLAG+
		budget.LastModificationDate.String())
}

func (s SmtpSenderBasic) SendBudgetArchive(budget models.Budget) error {
	return s.SendBudgetMessage(budget, models.BUDGET_ARCHIVE_FLAG+
		"["+archiveTimestamp()+"]")
}

func archiveTimestamp() string {
	return time.Now().Format("20060102150405")
}
