package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "guolmel",
	Short: "Simple budget management based on a webmail.",
	Long:  `It stands for "Good old mail".`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Guolmel.")
	},
}

func init() {}
