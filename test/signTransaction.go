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
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
    }

//Construct fromAddress by privatekey
//GNC7d4ed9084A364424d1087d26C4Ef092EAfB0b395(have 100000000 GNC)
    privateKey, err := crypto.HexToECDSA("a59bc058eb76eea5b64f1e55a803aa0968efda8a943f8f7eb835a6df9ac3a835")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(fromAddress.Hex())
//Construct toAddress
toAddress := common.HexToAddress("GNC30095Bb2A16CC8f4b897F511D2B62Fb8a0c2F0ec")

        
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