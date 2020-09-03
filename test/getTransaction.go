package main

import (
	"context"
	"log"
	// "math/big"
	"encoding/hex"
	"github.com/gnc-project/GNC-go/common"
	"github.com/gnc-project/GNC-go/core/types"
	"github.com/gnc-project/GNC-go/ethclient" 
)
func main() {
    client, err := ethclient.Dial("http://localhost:1234")
    if err != nil {
        log.Fatal(err)
	}
	var transaction *types.Transaction
	
	TransactionHash := common.HexToHash("0x05813c1580e2ba4f6d54264d82b204677b222b7a4da006eba30d2be65d572f51")
	transaction,isPending,err:=client.TransactionByHash(context.Background(),TransactionHash)
	if err != nil {
        log.Fatal(err)
	}
	if isPending==false{
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
	}
//get "From" must change msg type
	msg, err := transaction.AsMessage(types.NewEIP155Signer(chainID))
	if err != nil {
		log.Fatal("getFromErr",err)
	}

	log.Println("交易Hash:",transaction.Hash().Hex()) 
	log.Println("From:",msg.From().Hex())   
	log.Println("To:",transaction.To().Hex())      
	log.Println("Value:",transaction.Value().String())    
	log.Println("Gas:",transaction.Gas())               
	log.Println("GasPrice:",transaction.GasPrice().Uint64()) 
	log.Println("Nonce:",transaction.Nonce())             
	log.Println("Data:",hex.EncodeToString(transaction.Data()))
}
} 