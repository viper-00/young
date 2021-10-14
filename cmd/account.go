package cmd

import (
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var account = &cobra.Command{
	Use:   "account",
	Short: "Manage ethereum account",
	Run: func(cmd *cobra.Command, args []string) {
		// operation
		op, err := cmd.Flags().GetString("operation")
		if op == "" || err != nil {
			fmt.Println("please use flag to setup an operation")
			os.Exit(1)
		}
		switch op {
		case "balance":
			account := common.HexToAddress(Cfg.PrivateKey)
			balance, err := Cfg.ClientAddress.BalanceAt(cmd.Context(), account, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Eth wei => Eth value
			fbalance := new(big.Float)
			fbalance.SetString(balance.String())
			ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Println(ethValue)
		case "check":
			check, err := cmd.Flags().GetString("check")
			if check == "" || err != nil {
				fmt.Println("please use flag to setup an check")
				os.Exit(1)
			}
			address := common.HexToAddress(check)
			bytecode, err := Cfg.ClientAddress.CodeAt(cmd.Context(), address, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			isContract := len(bytecode) > 0
			if isContract {
				fmt.Println("This address is a Ethereum Contract")
			} else {
				fmt.Println("This address is a Ethereum Account")
			}
		default:
			fmt.Println("Use flag to setup an operation is unvalid")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(account)
	account.Flags().StringP("operation", "o", "", "account, example: balance | check")
	account.Flags().StringP("check", "c", "", "checkout address whether account or smart contract")

}
