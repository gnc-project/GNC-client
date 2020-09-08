package main

import (
	"fmt"
	"log"
    "crypto/ecdsa"
	"github.com/gnc-project/GNC-go/crypto"
    "github.com/gnc-project/GNC-go/common"
)


func main(){
	privateKey, err:=crypto.GenerateKey()
	if err!=nil{
		log.Fatal(err)
	}
	privateKeyString:=common.Bytes2Hex(privateKey.D.Bytes())
	fmt.Println("privateKey:",privateKeyString)
	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    address := crypto.PubkeyToAddress(*publicKeyECDSA)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("address:",address.Hex())
}