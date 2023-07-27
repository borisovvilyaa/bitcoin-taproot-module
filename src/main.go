package main

import (
	"bit/src/wallet_module"
	"log"
)

type AddressTaproot struct {
	wallets        []*wallet_module.Wallet
	AddressTaproot byte
}

func addWalletsToAddressTaproot(addressTaproot *AddressTaproot, wallets ...*wallet_module.Wallet) {
	addressTaproot.wallets = append(addressTaproot.wallets, wallets...)
}

func main() {
	// Create an instance of the AddressTaproot struct.

	// Generate some Taproot wallets.
	wallet1, err := wallet_module.GenerateWalletLegacy()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}
	wallet2, err := wallet_module.GenerateWalletLegacy()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}
	wallet3, err := wallet_module.GenerateWalletLegacy()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}
	wallet4, err := wallet_module.GenerateWalletLegacy()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}

	// Add wallets to the AddressTaproot struct using the function.
	taproot := wallet_module.GenerateTaproot(wallet1, wallet2, wallet3, wallet4)
	taproot.GetInfo()
}
