package cmd

import "github.com/spf13/cobra"

var query = &cobra.Command{
	Use:   "query",
	Short: "query network data",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(query)
}
