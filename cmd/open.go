package cmd

import (
	"errors"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/magleff/guolmel/mail/smtp"
	"github.com/magleff/guolmel/models"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"time"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a budget",
	Long:  `Open a budget.`,
	RunE:  Open,
}

func Open(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("You should provide an initial balance.")
	}

	initialBalance, errParsingAmount := decimal.NewFromString(args[0])

	if errParsingAmount != nil {
		return errParsingAmount
	}

	currentBudget, _ := imap.CurrentBudget()

	if currentBudget != nil {
		return errors.New("There is already an opened budget.")
	}

	budget := models.Budget{
		StartDate:            time.Now(),
		LastModificationDate: time.Now(),
		InitialBalance:       initialBalance,
	}

	budget.ComputeInformation()

	return new(smtp.SmtpSenderBasic).SendBudgetState(budget)
}

func init() {
	RootCmd.AddCommand(openCmd)
}
