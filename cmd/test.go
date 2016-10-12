package cmd

import (
	"errors"
	"fmt"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/magleff/guolmel/mail/smtp"
	"github.com/spf13/cobra"
	"os"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the configuration",
	Long:  `Test the configuration.`,
	RunE:  Test,
}

func Test(cmd *cobra.Command, args []string) error {
	envVarsToTest := []string{imap.SERVER, imap.USER, imap.PASSWORD,
		smtp.SERVER, smtp.USER, smtp.PASSWORD}

	for _, envVar := range envVarsToTest {
		if os.Getenv(envVar) == "" {
			return errors.New("Environment variable " + envVar + " not set.")
		}
	}

	mailSender := new(smtp.SmtpSenderBasic)
	errSmtp := mailSender.SendTestMail()

	if errSmtp != nil {
		return errSmtp
	}

	dialer := new(imap.ImapDialerBasic)
	_, errImap := dialer.DialTLS()

	if errImap != nil {
		return errImap
	}

	fmt.Println("Everything is ok.")

	return nil
}

func init() {
	RootCmd.AddCommand(testCmd)
}
