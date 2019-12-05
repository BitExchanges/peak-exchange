package erc20

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

type Transfer struct {
	client          *ethclient.Client
	contractAddress common.Address
}

// 初始化主网连接
func InitClient(url, contractAddress string) *Transfer {
	rpcDial, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcDial)
	return &Transfer{client: client, contractAddress: common.HexToAddress(contractAddress)}
}

// 外部转账
func (t *Transfer) OutTransfer(fromAddress common.Address, toAddress string, amount int64) *types.Transaction {
	//计算nonce值
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//1eth =1000000000000000000 wei
	value := big.NewInt(amount * 1000000000000000000)
	//一次打包可以使用的gas上限
	gasLimit := uint64(21000)

	//由于不确定gas费用，取平均值
	gasPrice, err := t.client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)
	return tx
}

// 交易签名并发送
func (t *Transfer) SignTx(tx *types.Transaction, privateKeyStr string) *types.Transaction {
	//字符串私钥转换为 ecdsa
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	//获取网络通道ID
	chainId, err := t.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//对transaction做交易签名
	signedTX, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedTX
}
