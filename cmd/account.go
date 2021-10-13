package cmd

import "github.com/spf13/cobra"

var account = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(account)
}
