package cmd

import (
	"errors"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/magleff/guolmel/mail/smtp"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Closes and archives the current budget.",
	Long:  `Closes and archives the current budget.`,
	RunE:  Close,
}

func Close(cmd *cobra.Command, args []string) error {
	currentBudget, _ := imap.CurrentBudget()

	if currentBudget == nil {
		return errors.New("There is not any opened budget.")
	}

	new(smtp.SmtpSenderBasic).SendBudgetArchive(*currentBudget)

	imap.PurgeBudgetStates()

	return nil
}

func init() {
	RootCmd.AddCommand(closeCmd)
}
