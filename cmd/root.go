package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Young",
		Short: "Young builds your ethereum client",
		Long: `Young is the main command, used to connect ethereum network.

Complete documentation is available at https://github.com/zhong-my/young`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {

}
