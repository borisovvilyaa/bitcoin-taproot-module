package wallet_module

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// Wallet represents a structure with a key pair and a bitcoin wallet address.
type Wallet struct {
	privateKey *btcec.PrivateKey
	publicKey  *btcec.PublicKey
	address    btcutil.Address
}

// GenerateWalletLegacy generates a new key pair and Bitcoin wallet address (for legacy logic).
func GenerateWalletLegacy() (*Wallet, error) {
	// Create a new private key.
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	// Obtain the public part of the key.
	publicKey := privateKey.PubKey()

	// Get the Bitcoin wallet address from the uncompressed public key.
	address, err := btcutil.NewAddressPubKey(publicKey.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKey,
		address:    address,
	}, nil
}

// GenerateWalletTaproot generates a new key pair and Bitcoin wallet address using Taproot.
func GenerateWalletTaproot() (*Wallet, error) {
	// Create a new private key.
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	// Obtain the public part of the key.
	publicKey := privateKey.PubKey()

	// Create a SegWit script from the compressed public key.
	script, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(publicKey.SerializeCompressed()), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// Generate a Taproot address from the SegWit script.
	address, err := btcutil.NewAddressTaproot([]byte(script.EncodeAddress()), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKey,
		address:    address,
	}, nil
}

// ToString returns a string representation of the key pair and address.
func (w *Wallet) ToString() string {
	privateWIF, err := btcutil.NewWIF(w.privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "Error: Unable to generate WIF for private key."
	}

	str := fmt.Sprintf("Bitcoin Private Key (WIF): %s\n", privateWIF.String())
	str += fmt.Sprintf("Bitcoin Public Key: %x\n", w.publicKey.SerializeCompressed())
	str += fmt.Sprintf("Bitcoin Address: %s", w.address.EncodeAddress())

	return str
}
