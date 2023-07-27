package wallet_module

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

type AddressTaproot struct {
	wallets        []*Wallet
	hashedPk       []string
	MarkleeTree    string
	AddressTaproot string
}

func addWalletsToAddressTaproot(addressTaproot *AddressTaproot, wallets ...*Wallet) {
	addressTaproot.wallets = append(addressTaproot.wallets, wallets...)

}

func hashAddressSHA256(addressTaproot *AddressTaproot, wallets ...*Wallet) {
	for _, wallet := range addressTaproot.wallets {
		inputBytes := []byte(wallet.publicKey.SerializeCompressed())

		// Создаем объект хеша SHA-256
		sha256Hash := sha256.New()

		// Записываем входные данные в хеш
		sha256Hash.Write(inputBytes)

		// Получаем хеш в виде байтового среза ([]byte)
		hashedBytes := sha256Hash.Sum(nil)

		// Преобразуем хеш из байтового среза в строку шестнадцатеричных символов
		hashedString := hex.EncodeToString(hashedBytes)
		addressTaproot.hashedPk = append(addressTaproot.hashedPk, hashedString)
	}

}
func calculateHash(data1, data2 string) string {
	hash := sha256.New()
	hash.Write([]byte(data1 + data2))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Функция для построения дерева Меркла из списка хэшей
func buildMerkleTree(hashes []string) []string {
	if len(hashes) == 0 {
		return nil
	}

	// Если в списке хэшей только один элемент, возвращаем его
	if len(hashes) == 1 {
		return []string{hashes[0]}
	}

	var merkleTree []string

	// Если количество хэшей нечетное, дублируем последний хэш
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// Построение дерева Меркла
	for i := 0; i < len(hashes); i += 2 {
		// Вычисляем хэш для пары элементов
		hash := calculateHash(hashes[i], hashes[i+1])
		merkleTree = append(merkleTree, hash)
	}

	// Рекурсивный вызов функции для построения следующего уровня дерева
	return buildMerkleTree(merkleTree)
}

func bytesToBech32Address(data []byte) string {
	hrp := "bc"
	convertedData, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return ""
	}

	bech32Address, err := bech32.Encode(hrp, convertedData)
	if err != nil {
		return ""
	}

	return bech32Address
}

func GenerateTaproot(wallets ...*Wallet) AddressTaproot {
	addressTaproot := AddressTaproot{}
	addWalletsToAddressTaproot(&addressTaproot, wallets...)
	hashAddressSHA256(&addressTaproot, wallets...)
	addressTaproot.MarkleeTree = buildMerkleTree(addressTaproot.hashedPk)[0]
	addressTaproot.AddressTaproot = bytesToBech32Address([]byte(addressTaproot.MarkleeTree))
	return addressTaproot
}

func (addressTaproot *AddressTaproot) GetInfo() {
	fmt.Println(addressTaproot)
	for n, wallwallet := range addressTaproot.wallets {
		fmt.Println(wallwallet.ToString())
		fmt.Println("Hashed Double PK:", addressTaproot.hashedPk[n])
		fmt.Println()

	}
	fmt.Println("Marklee Tree", addressTaproot.MarkleeTree)
	fmt.Println("bech32 T2TP", addressTaproot.AddressTaproot)
}
