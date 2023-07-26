package main

import (
	"fmt"

	wallet "bit/src/wallet_module"
)

func main() {
	wallet, err := wallet.GenerateWalletLegacy()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(wallet.ToString())
}
