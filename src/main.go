package main

import (
	"bit/src/wallet_module"
	"fmt"
	"log"
)

func main() {
	// Generate a new Taproot wallet.
	wallet, err := wallet_module.GenerateWalletLegacy()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}

	// Display wallet information.
	fmt.Println("Wallet Information:")
	fmt.Println(wallet.ToString())
}
