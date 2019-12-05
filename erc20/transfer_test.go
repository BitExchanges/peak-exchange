package erc20

import (
	"context"
	"testing"
)

func TestTransfer_OutTransfer(t *testing.T) {
	tran := InitClient("http://127.0.0.1:7545", "0x812F6Fe50b7189CB2f6031846cf9Ef0f395a1C20")
	tx := tran.OutTransfer(tran.contractAddress, "0x914e0cfA008293cc14634F5C0a4284d9242f2077", 2)
	signedTX := tran.SignTx(tx, "26968f6c1cc0ddc53897c4feb1c1256b3e93575143c1de4ae7558d1ea2bef61c")
	tran.client.SendTransaction(context.Background(), signedTX)
}
