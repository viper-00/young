package cmd

import (
	"fmt"
	"os"
	"young/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Young",
		Short: "Young builds your ethereum client",
		Long: `Young is the main command, used to connect ethereum network.

Complete documentation is available at https://github.com/zhong-my/young`,
	}
	Cfg       config.Config
	initState bool
)

func init() {
	cobra.OnInitialize(initConfig, initNetWork, initDone)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if !initState {
		// Set the file name of the configurations file
		viper.SetConfigName("config")

		// Set the file extension of the configurations file
		viper.SetConfigType("yaml")

		// Set the path to look for the configurations file
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
	}
}

func initNetWork() {
	address := viper.GetString("client_address")
	key := viper.GetString("private_key")
	client, err := ethclient.Dial(address)
	if err != nil {
		panic(err)
	}
	Cfg.ClientAddress = client
	Cfg.PrivateKey = key
}

func initDone() {
	initState = true
}
