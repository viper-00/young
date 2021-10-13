package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print version numbers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Young")
		fmt.Println("Version: 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(version)
}
