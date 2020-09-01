package main

import (
    "context"
    "log"
    "github.com/gnc-project/GNC-go/common"
    "github.com/gnc-project/GNC-go/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

	Address := common.HexToAddress("0xf948a50e7Deae42c81e205CE73220a28bEDa71d7")

	balance,err:=client.BalanceAt(context.Background(),Address,nil)
	if err != nil {
        log.Fatal(err)
    }
    log.Println("balance===>",balance)
}