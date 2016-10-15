package cmd

import (
	"fmt"
	"github.com/magleff/guolmel/mail/imap"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the current status",
	Long:  `Gives the current status.`,
	RunE:  Status,
}

func Status(cmd *cobra.Command, args []string) error {
	fmt.Println(imap.CurrentBudget())
	return nil
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
