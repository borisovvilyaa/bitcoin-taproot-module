package wallet_module

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/btcsuite/btcd/chaincfg"
)

// ToString returns a string representation of the key pair and addresses.
func (w *Wallet) ToString() string {
	privateWIF, err := btcutil.NewWIF(w.privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return "Error: Unable to generate WIF for private key."
	}

	str := fmt.Sprintf("Bitcoin Private Key (WIF): %s\n", privateWIF.String())
	str += fmt.Sprintf("Bitcoin Public Key: %x\n", w.publicKey.SerializeCompressed())
	str += fmt.Sprintf("Bitcoin P2WPKH: %s\n", w.SegWitAddress.EncodeAddress())
	str += fmt.Sprintf("Bitcoin P2TR: %s", w.taprootAddress)

	return str
}

// Wallet represents a structure with a key pair and bitcoin addresses (legacy and taproot).
type Wallet struct {
	privateKey     *btcec.PrivateKey
	publicKey      *btcec.PublicKey
	SegWitAddress  btcutil.Address
	taprootAddress string // Taproot address in Bech32 format
}

// GenerateWallet generates a new key pair and Bitcoin wallet addresses (SegWigt and taproot).
func GenerateWallet() (*Wallet, error) {
	// Create a new private key.
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	// Obtain the public part of the key.
	publicKey := privateKey.PubKey()

	// Get the Bitcoin legacy address from the compressed public key (P2WPKH - Pay to Witness Public Key Hash).
	legacyAddress, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(publicKey.SerializeCompressed()), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// Generate the Taproot address (P2TR - Pay to Taproot).
	taprootScript, err := bech32.ConvertBits(publicKey.SerializeCompressed(), 8, 5, true)
	if err != nil {
		return nil, err
	}
	taprootAddress, err := bech32.Encode("bc", taprootScript)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		privateKey:     privateKey,
		publicKey:      publicKey,
		SegWitAddress:  legacyAddress,
		taprootAddress: taprootAddress,
	}, nil
}

// getPubKey returns the public key of the wallet.
func (w *Wallet) GetPubKey() *btcec.PublicKey {
	return w.publicKey
}

// getPrivateKey returns the private key of the wallet.
func (w *Wallet) GetPrivateKey() *btcec.PrivateKey {
	return w.privateKey
}

// getPrivateKey returns the private key of the wallet.
func (w *Wallet) GetP2WPKH() btcutil.Address {
	return w.SegWitAddress
}

// getPrivateKey returns the private key of the wallet.
func (w *Wallet) GetP2TR() string {
	return w.taprootAddress
}
