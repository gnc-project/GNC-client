package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    // "strconv"
    "math"
    "github.com/shopspring/decimal"
    "github.com/gnc-project/GNC-go/common"
    "github.com/gnc-project/GNC-go/core/types"
    "github.com/gnc-project/GNC-go/crypto"
    "github.com/gnc-project/GNC-go/ethclient"
)
func main() {
    //Connect node
    client, err := ethclient.Dial("http://120.79.174.236:8545")
    if err != nil {
        log.Fatal(err)
    }

//Construct fromAddress by privatekey
//GNC30095Bb2A16CC8f4b897F511D2B62Fb8a0c2F0ec
    // privateKey, err := crypto.HexToECDSA("b77de610fb69f929f9ce38e07bc003bb8dfffc9024c0af0da26ab2d0a052492e")
//0x0108aE381335Bba1F5a3293D501947D6174de367
    // privateKey, err := crypto.HexToECDSA("6ab0638768979e4a551a2c81b90c943cb12e07819bee721be74aaf481919bb2b")
    privateKey, err := crypto.HexToECDSA("b7b5f09a0a7147a64e17f51d3f220211cf3a6d0d647d81eaf970d5f3e0a93e33")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    fmt.Println(fromAddress.Hex())
    return 
    if err != nil {
        log.Fatal(err)
    }
//Construct toAddress
toAddress := common.HexToAddress("30095Bb2A16CC8f4b897F511D2B62Fb8a0c2F0ec")

        
//value
value := decimal.NewFromFloat(0.1)//this is value you want to send

decimals := decimal.NewFromFloat(math.Pow10(18))
amount:=value.Mul(decimals)//Authentic value 
    
//gasPrice
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    gas:=uint64(21000)
//nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//data   
    data:=[]byte("")
//Construct transaction
    tx := types.NewTransaction(nonce, toAddress, amount.BigInt(),gas, gasPrice, data)
//Inquire chainID
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }
//Sign transaction 
    var signedTx *types.Transaction
    signedTx, err = types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID.Int64())), privateKey)
    if err != nil {
        log.Fatal(err)    
    }
//send signatureTx 
    err = client.SendRawTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("tx Hash: %v\n", signedTx.Hash().Hex())
    log.Println("Waiting for the transaction, about 4 minutes...")
//wait TX
	for {
    tx, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
    if err != nil {
        log.Fatal(err)
    }
    if isPending==false{
         fmt.Println("transaction is successful!!")
		 receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
        }
        if receipt.Status==0{
            log.Fatal( "Error: Transaction has been reverted by the EVM")
        }
		fmt.Printf("receipt.Status:%v\n",receipt.Status)
		return 
    }
   }
}