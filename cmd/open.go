package cmd

import (
	"github.com/magleff/guolmel/imap"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a budget",
	Long:  `Open a budget.`,
	RunE:  Open,
}

func Open(cmd *cobra.Command, args []string) error {
	dialer := new(imap.ImapDialerBasic)
	_, err := dialer.DialTLS()

	if err != nil {
		return err
	}

	return nil
}

func init() {
	RootCmd.AddCommand(openCmd)
}
