package wallet_module

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

type Wallet struct {
	Address *btcutil.AddressTaproot
	Type    string
}

// func (wallet *KeyPair) GenerateWalletLegacy(Wallet) {
// 	address, err := btcutil.NewAddressPubKey(publicKey.SerializeUncompressed(), &chaincfg.MainNetParams)

// }

func (wallet *KeyPair) GenerateWalletTaproot() (*Wallet, error) {
	taprootAddress, err := btcutil.NewAddressTaproot(wallet.publicKey.SerializeCompressed()[0:32], &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	return &Wallet{
		Address: taprootAddress,
		Type:    "T2PR",
	}, err
}
func getAddressBalance(address *btcutil.AddressTaproot) (int, error) {
	baseURL := "https://blockchain.info/q/addressbalance/"
	url := baseURL + address.EncodeAddress()
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.Atoi(string(body))
	if err != nil {
		return 0, err
	}

	return balance, nil
}
func (w *Wallet) ToString() string {
	str := fmt.Sprintf("Bitcoin Public Key: %v\n", w.Address)
	str += fmt.Sprintf("Type: %s\n", w.Type)
	balance, err := getAddressBalance(w.Address)
	if err != nil {
		fmt.Println("Error:", err)
	}

	str += fmt.Sprintf("Баланс адреса: %d BTC\n", balance)

	// str += fmt.Sprintf("Bitcoin Address: %s", w.address.EncodeAddress())

	return str
}
