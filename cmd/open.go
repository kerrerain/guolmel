package cmd

import (
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a budget",
	Long:  `Open a budget.`,
	RunE:  Open,
}

func Open(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	RootCmd.AddCommand(openCmd)
}
