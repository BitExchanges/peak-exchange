package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

var EthClient *ethclient.Client

//初始化以太坊客户端
func InitEthClient(url string) {
	rpcDial, err := rpc.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("以太坊客户端连接成功")
	client := ethclient.NewClient(rpcDial)
	EthClient = client
}

func init() {
	InitEthClient("http://127.0.0.1:7545")
}
