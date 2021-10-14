package cmd

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var query = &cobra.Command{
	Use:   "query",
	Short: "query ethereum network data",
	Run: func(cmd *cobra.Command, args []string) {
		// operation
		op, err := cmd.Flags().GetString("operation")
		if op == "" || err != nil {
			fmt.Println("please use flag to setup an operation")
			os.Exit(1)
		}

		switch op {
		case "block_by_hash":
			hash, err := cmd.Flags().GetString("hash")
			if hash == "" || err != nil {
				fmt.Println("please use flag to setup a hash")
				os.Exit(1)
			}

			block, err := Cfg.ClientAddress.BlockByHash(cmd.Context(), common.HexToHash(hash))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Block result ->")
			fmt.Printf("Block Number: %d\n", block.Number().Uint64())
			fmt.Printf("Block Time: %d\n", block.Time())
			fmt.Printf("Block Difficulty: %d\n", block.Difficulty().Uint64())
			fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
			fmt.Printf("Block Transactions: %d\n", len(block.Transactions()))
		case "block_by_number":
			number, err := cmd.Flags().GetInt("number")
			if number == 0 || err != nil {
				fmt.Println("please use flag to setup a number")
				os.Exit(1)
			}
			block, err := Cfg.ClientAddress.BlockByNumber(cmd.Context(), big.NewInt(int64(number)))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Block result ->")
			fmt.Printf("Block Number: %d\n", block.Number().Uint64())
			fmt.Printf("Block Time: %d\n", block.Time())
			fmt.Printf("Block Difficulty: %d\n", block.Difficulty().Uint64())
			fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
			fmt.Printf("Block Transactions: %d\n", len(block.Transactions()))

		default:
			fmt.Println("Use flag to setup an operation is unvalid")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(query)
	query.Flags().StringP("operation", "o", "", "query that what do you want, example: block_by_hash | block_by_number")
	query.Flags().StringP("hash", "s", "", "hash from block")
	query.Flags().IntP("number", "n", 0, "number from block")
}
