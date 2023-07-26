package wallet_module

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// Wallet представляет структуру с ключевой парой и адресом биткоин-кошелька.
type Wallet struct {
	privateKey *btcec.PrivateKey
	publicKey  *btcec.PublicKey
	address    btcutil.Address
}

// GenerateWallet генерирует новую ключевую пару и адрес биткоин-кошелька.
func GenerateWalletLegacy() (*Wallet, error) {
	// Создаем новый ключ.
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	// Получаем публичную часть ключа.
	publicKey := privateKey.PubKey()

	// Получаем адрес биткоин-кошелька из публичного ключа.
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
