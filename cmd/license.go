package cmd

import "github.com/spf13/cobra"

var license = &cobra.Command{
	Use:   "license",
	Short: "Display license information",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(license)
}
