package cmd

import (
	"fmt"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/magleff/guolmel/models"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the current status",
	Long:  `Gives the current status.`,
	RunE:  Status,
}

func Status(cmd *cobra.Command, args []string) error {
	budget, err := imap.CurrentBudget()

	if err != nil {
		return err
	}

	displayBudgetInformation(budget)

	return nil
}

func displayBudgetInformation(budget *models.Budget) {
	fmt.Printf("\n")
	fmt.Println("Created on", budget.StartDate)
	fmt.Println("Initial balance", budget.InitialBalance)
	fmt.Println("Total earnings", budget.TotalEarnings)
	fmt.Println("Total expenses", budget.TotalExpenses)
	fmt.Println("Total unchecked expenses", budget.TotalUncheckedExpenses)
	fmt.Println("Balance", budget.CurrentBalance.String(), "("+budget.Difference.String()+")")
	displayBudgetExpenses(budget)
}

func displayBudgetExpenses(budget *models.Budget) {
	fmt.Printf("\n")
	fmt.Printf("%-2s|%-6s|%-16s|%-30s|%-16s\n", "C", "Index", "Amount", "Description", "Date")
	fmt.Printf("\n")
	for index, entry := range budget.Expenses {
		fmt.Printf("%-2s|%-6v|%-16v|%-30s|%-16s\n",
			printChecked(entry.Checked),
			index,
			entry.Amount,
			entry.Description,
			entry.Date.Format("2006-01-02"))
	}
	fmt.Printf("\n")
}

func printChecked(checked bool) string {
	return printBoolean(checked, "X")
}

func printBoolean(boolean bool, char string) string {
	str := ""
	if boolean {
		str = char
	}
	return str
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
