package main

import (
    "context"
    "log"
    "github.com/gnc-project/GNC-go/common"
    "github.com/gnc-project/GNC-go/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
    }

	Address := common.HexToAddress("GNC7d4ed9084A364424d1087d26C4Ef092EAfB0b395")

	balance,err:=client.BalanceAt(context.Background(),Address,nil)
	if err != nil {
        log.Fatal(err)
    }
    log.Println("balance===>",balance)
}