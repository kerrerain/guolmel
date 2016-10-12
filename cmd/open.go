package cmd

import (
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
	budget := models.Budget{
		StartDate:            time.Now(),
		LastModificationDate: time.Now(),
		InitialBalance:       decimal.NewFromFloat(0.00),
	}

	mailSender := new(smtp.SmtpSenderBasic)
	return mailSender.SendBudgetState(budget)
}

func init() {
	RootCmd.AddCommand(openCmd)
}
