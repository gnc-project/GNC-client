package main

import (
	"context"
	"log"
	// "math/big"

	"github.com/gnc-project/GNC-go/common"
	"github.com/gnc-project/GNC-go/core/types"
	"github.com/gnc-project/GNC-go/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
	}
	var block *types.Block
	blockHash := common.HexToHash("0x18c8c36ac3c285d7b276e59b1988d0632aec58ee7f70faa17cfe74de0c5484b5")
	block,err=client.BlockByHash(context.Background(),blockHash)
//or ByNumber
	// block,err=client.BlockByNumber(context.Background(),big.NewInt(25591))

	if err != nil {
        log.Fatal(err)
	}
	header:=block.Header()
	log.Println("blockHash===>",block.Hash().Hex())
	log.Println("blockNumber===>",block.Number())
	log.Println("blockHash:",header.Hash().Hex()) 
	log.Println("parentHash:",header.ParentHash.Hex())
	log.Println("blockNumber:",block.Number().Uint64())    
	log.Println("blockCoinbase:",block.Coinbase().Hex())    
	log.Println("timestamp:",block.Time())    
	log.Println("gasLimit:",block.GasLimit())    
	log.Println("size:",block.Size()) 
}