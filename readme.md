## GNC-go

```txt
Examples are in the test folder
```

### Connect node

```go
    client, err := ethclient.Dial("http://localhost:1234")
```
### Send transaction (signature mode)

#### Construct fromAddress by privatekey
```go
//0xf5403E4F120901407eF221E2419583D1F3556953
    privateKey, err := crypto.HexToECDSA("b77de610fb69f929f9ce38e07bc003bb8dfffc9024c0af0da26ab2d0a052492e")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
```
#### Construct toAddress
```go
   toAddress := common.HexToAddress("GNC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")
```
#### set value
```go
    value := decimal.NewFromFloat(0.1)//this is value you want to send

    decimals := decimal.NewFromFloat(math.Pow10(18))
    amount:=value.Mul(decimals)//Authentic value 
```
#### set gasPrice and gaslimit
```go 
//gasPrice
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
//gaslimit
    gas:=uint64(21000)
```
#### set nonce 
```go
//nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
```
#### Construct transaction
```go
 //"github.com/gnc-project/GNC-go/core/types"
    tx := types.NewTransaction(nonce, toAddress, amount.BigInt(),gas, gasPrice, data)
```
#### Inquire chainID
```go
   chainID, err := client.NetworkID(context.Background())
```
#### Sign transaction
```go
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID.Int64())), privateKey)
```
#### send signatureTx
```go
//"github.com/gnc-project/GNC-go/core/types"
    err = client.SendRawTransaction(context.Background(), signedTx)
```
### Send transaction (unlock mode)
```go
    fromAddress := common.HexToAddress("GNC0d8c6aba421723b3bce849c70c06592f696e4399")
    toAddress := common.HexToAddress("GNC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")

    //Inquire nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
        if err != nil {
            log.Fatal(err)
        }
    //Construct transaction
    tx:=types.SendTxArgs{
        From:       fromAddress,
        To :        toAddress,
        Gas:        hexutil.Uint64(21000),
        GasPrice:   hexutil.Big(*big.NewInt(6000000000)),
        Value:      hexutil.Big(*big.NewInt(1)),
        Nonce:      hexutil.Uint64(nonce),
 
        Data:      hexutil.Bytes([]byte{}),
        Type:      0,
    }
    //send transaction
//"github.com/gnc-project/GNC-go/core/types"
        hash,err:=client.SendTransaction(context.Background(),tx)
        if err != nil {
            log.Fatal(err)
        }
```
### Inquire transaction
```go
	transaction,isPending,err:=client.TransactionByHash(context.Background(),TransactionHash)
	if err != nil {
        log.Fatal(err)
    }
```
### Inquire Balance
```go
	Address := common.HexToAddress("GNC86056D210eA7Bc23337aCaBE96dE275E584a67ce")
	balance,err:=client.BalanceAt(context.Background(),Address,nil)
	if err != nil {
        log.Fatal(err)
    }
```
### Inquire block
```go
	blockHash := common.HexToHash("0x18c8c36ac3c285d7b276e59b1988d0632aec58ee7f70faa17cfe74de0c5484b5")
    block,err=client.BlockByHash(context.Background(),blockHash)
    //or ByNumber
	// block,err=client.BlockByNumber(context.Background(),big.NewInt(25591))
```
