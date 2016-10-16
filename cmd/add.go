package cmd

import (
	"errors"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/magleff/guolmel/mail/smtp"
	"github.com/magleff/guolmel/models"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an expense",
	Long:  `Adds an expense.`,
	RunE:  Add,
}

func parseString(amount string) (decimal.Decimal, error) {
	amount = strings.Replace(amount, ",", ".", 1)
	amountDecimal, err := decimal.NewFromString(amount)

	if err != nil {
		return decimal.NewFromFloat(-1),
			errors.New("The given string " + amount + " is not a number.")
	}

	return amountDecimal.Truncate(2), nil
}

func ParseStringAmount(amount string) (decimal.Decimal, error) {
	parsedAmount, err := parseString(amount)

	if err != nil {
		return parsedAmount, err
	}

	if !strings.Contains(amount, "+") {
		parsedAmount = parsedAmount.Mul(decimal.NewFromFloat(-1))
	}

	return parsedAmount, nil
}

func ParseExpenseFromArguments(args []string) (*models.Expense, error) {
	var description string

	amount, err := ParseStringAmount(args[0])

	if err != nil {
		return nil, err
	}

	if len(args) > 1 {
		description = args[1]
	} else {
		description = ""
	}

	return &models.Expense{
		Amount:      amount,
		Description: description,
	}, nil
}

func Add(cmd *cobra.Command, args []string) error {
	currentBudget, _ := imap.CurrentBudget()

	if currentBudget == nil {
		return errors.New("There is not any opened budget.")
	}

	expense, err := ParseExpenseFromArguments(args)

	if err != nil {
		return err
	}

	expense.Date = time.Now()
	expense.Checked = false

	currentBudget.AddExpense(*expense)

	return new(smtp.SmtpSenderBasic).SendBudgetState(*currentBudget)
}

func init() {
	RootCmd.AddCommand(addCmd)
}
