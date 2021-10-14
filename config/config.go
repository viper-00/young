package config

import "github.com/ethereum/go-ethereum/ethclient"

// main config
type Config struct {
	ClientAddress *ethclient.Client
	PrivateKey    string
	Gas           string
	GasPrice      string
	Nonce         string
	From          string
	To            string
}
