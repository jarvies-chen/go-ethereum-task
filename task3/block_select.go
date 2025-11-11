package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/types"
)

func queryBlockInfo() {
	// 创建 RPC 客户端
	rpcClient := client.NewClient("https://api.devnet.solana.com")

	// 获取最新区块高度
	slot, err := rpcClient.GetSlot(context.Background())
	if err != nil {
		log.Printf("获取区块高度失败: %v", err)
		return
	}

	fmt.Printf("当前区块高度 (Slot): %d\n", slot)

	// 获取区块信息
	block, err := rpcClient.GetBlock(context.Background(), slot)
	if err != nil {
		log.Printf("获取区块信息失败: %v", err)
		return
	}

	fmt.Printf("Block Hash: %s\n", block.Blockhash)
	fmt.Printf("Previous Block Hash: %s\n", block.PreviousBlockhash)
	fmt.Printf("Transactions Count: %d\n", len(block.Transactions))
}

func getAccountBalance() {
	rpcClient := client.NewClient("https://api.devnet.solana.com")

	//创建账户
	account := types.NewAccount()
	fmt.Printf("Account public key: %s\n", account.PublicKey.ToBase58())

	// 查询余额
	balance, err := rpcClient.GetBalance(context.Background(), account.PublicKey.ToBase58())
	if err != nil {
		log.Printf("查询余额失败: %v", err)
		return
	}

	fmt.Printf("余额: %d ", balance)
}

func main() {
	queryBlockInfo()
	getAccountBalance()
}
