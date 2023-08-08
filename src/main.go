package main

import (
	"bit/src/wallet_module"
	"fmt"
)

func generateWalletTaproot() *wallet_module.Wallet {
	keyPair, err := wallet_module.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating", err)
	}
	taproot, err := keyPair.GenerateWalletTaproot()
	if err != nil {
		fmt.Println("Error generating", err)
	}
	// fmt.Println(taproot)
	return taproot
}
func main() {
	for i := 0; i < 10; i++ {
		Address := generateWalletTaproot()
		fmt.Println(Address.ToString())
	}

}
