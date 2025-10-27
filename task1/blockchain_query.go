package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func queryBlock(blockNumber int64) {
	//连接到 Sepolia 测试网络
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//查询指定区块号的区块信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return
	}

	// 输出区块信息
	fmt.Printf("区块号: %d\n", block.NumberU64())
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("父区块哈希: %s\n", block.ParentHash().Hex())
	fmt.Printf("时间戳: %d (%s)\n", block.Time(), time.Unix(int64(block.Time()), 0).Format(time.RFC3339))
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("难度: %s\n", block.Difficulty().String())
	fmt.Printf("Gas限制: %d\n", block.GasLimit())
	fmt.Printf("Gas使用量: %d\n", block.GasUsed())
	fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
	fmt.Printf("区块大小: %d bytes\n", block.Size())
}

func main() {
	// 查询指定区块号的区块信息
	queryBlock(9226231)
}
