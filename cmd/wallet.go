package cmd

import "github.com/spf13/cobra"

var wallet = &cobra.Command{
	Use:   "wallet",
	Short: "Manage Ethereum presale wallets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(wallet)
}
