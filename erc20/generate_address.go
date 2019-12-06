package erc20

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

type Erc20Address struct {
	privateKey string
	address    string
}

type Erc20AddressArray []Erc20Address

func GenerateAddress(count int) []Erc20Address {
	var erc20Address []Erc20Address
	for i := 0; i < count; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
			continue
		}
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		privateKeyString := hex.EncodeToString(privateKey.D.Bytes())

		erc20Address = append(erc20Address, Erc20Address{
			privateKey: privateKeyString,
			address:    address.Hex(),
		})
	}

	//for _, v := range erc20Address {
	//	fmt.Printf("私钥: %s  地址: %s\n", v.privateKey, v.address)
	//}
	return erc20Address
}

// 生成私钥和地址
func GenerateUserWallet() (key, add string) {
	//生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return hex.EncodeToString(privateKey.D.Bytes()), address.Hex()
}
