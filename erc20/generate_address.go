package erc20

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

type Erc20Address struct {
	privateKey string
	address    string
}

type Erc20AddressArray []Erc20Address

func GenerateAddress(count int) {
	var addressArray Erc20AddressArray
	for i := 0; i < count; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
			continue
		}
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		privateKeyString := hex.EncodeToString(privateKey.D.Bytes())
		addressArray = append(addressArray, Erc20Address{
			privateKey: privateKeyString,
			address:    address.Hex(),
		})
	}

	for _, v := range addressArray {

		fmt.Printf("私钥: %s  地址: %s\n", v.privateKey, v.address)
	}
}

// 生成私钥和地址
func GenerateUserAddress() (key, add string) {
	//生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return hex.EncodeToString(privateKey.D.Bytes()), address.Hex()
}
