package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/types"
)

func transaction() {
	// 创建 RPC 客户端
	rpcClient := client.NewClient("https://api.devnet.solana.com")
	// 创建发送方和接收方账户
	sender := types.NewAccount()
	receiver := types.NewAccount()

	fmt.Printf("发送方地址: %s\n", sender.PublicKey.ToBase58())
	fmt.Printf("接收方地址: %s\n", receiver.PublicKey.ToBase58())

	// 获取最新区块哈希
	response, err := rpcClient.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("获取区块哈希失败: %v", err)
	}

	// 创建转账指令
	instruction := system.Transfer(system.TransferParam{
		From:   sender.PublicKey,
		To:     receiver.PublicKey,
		Amount: 1_000_000, // 0.001 SOL (1 SOL = 1,000,000,000 lamports)
	})

	// 构建交易消息
	message := types.NewMessage(types.NewMessageParam{
		FeePayer:        sender.PublicKey,
		RecentBlockhash: response.Blockhash,
		Instructions:    []types.Instruction{instruction},
	})

	// 创建交易并签名
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: message,
		Signers: []types.Account{sender},
	})
	if err != nil {
		log.Fatalf("创建交易失败: %v", err)
	}

	// 发送交易
	txHash, err := rpcClient.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Printf("交易已发送! 交易哈希: %s\n", txHash)
	fmt.Printf("在浏览器查看: https://explorer.solana.com/tx/%s?cluster=devnet\n", txHash)

}

func main() {
	transaction()
}
