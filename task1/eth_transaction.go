package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendTransaction(privateKeyHex, toAddressHex string, amount float64) {
	//1. 连接到 Sepolia 测试网络
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//2. 将私钥字符串转换为私钥对象
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	//3. 获取发送方公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	//4. 通过公钥获取发送方地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//5. 获取 nonce（交易计数）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//6. 获取 gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//7. 将 ETH 金额转换为 wei
	value := new(big.Int)
	value.SetString(fmt.Sprintf("%.0f", amount*1e18), 10)

	//8. 创建交易
	toAddress := common.HexToAddress(toAddressHex)

	tx := types.NewTransaction(
		nonce,
		toAddress,
		value,
		uint64(21000),
		gasPrice,
		nil,
	)

	//9. 获取链 ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//10. 签名交易
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//11. 发送交易
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送！\n")
	fmt.Printf("交易哈希: %s\n", signTx.Hash().Hex())
	fmt.Printf("发送方地址: %s\n", fromAddress.Hex())
	fmt.Printf("接收方地址: %s\n", toAddress)
	fmt.Printf("转账金额: %f ETH\n", amount)
	fmt.Printf("Gas价格: %s Gwei\n", new(big.Float).Quo(new(big.Float).SetInt(gasPrice), big.NewFloat(1e9)).Text('f', 2))
}

func main() {
	// 发送 0.001 ETH 到指定地址
	privateKeyHex := "YOUR_PRIVATE_KEY_HERE" // 替换为你的私钥
	toAddress := "TO_ADDRESS"
	amount := 1.0 // 发送 1 ETH

	sendTransaction(privateKeyHex, toAddress, amount)
}
