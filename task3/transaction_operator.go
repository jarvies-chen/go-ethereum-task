package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

func main() {
	// 1. 创建 RPC 客户端连接到 Devnet
	rpcClient := client.NewClient(rpc.DevnetRPCEndpoint)

	// 2. 发送方
	sender := types.NewAccount()

	// 3. 接收方
	receiver := types.NewAccount()

	// 4. 定义转账金额 (单位: Lamports)
	lamports := uint64(1000000) // 0.001 SOL

	// 5. 获取最新区块哈希 (用于交易)
	blockhashResp, err := rpcClient.GetLatestBlockhash(context.TODO())
	if err != nil {
		log.Fatalf("Failed to get latest blockhash: %v", err)
	}
	recentBlockhash := blockhashResp.Blockhash
	fmt.Printf("Latest blockhash for transaction: %s\n", recentBlockhash)

	// 6. 构造转账指令
	transferInstruction := system.Transfer(system.TransferParam{
		From:   sender.PublicKey,
		To:     receiver.PublicKey,
		Amount: lamports,
	})

	// 7. 构建交易消息
	message := types.NewMessage(types.NewMessageParam{
		FeePayer:        sender.PublicKey,
		RecentBlockhash: recentBlockhash,
		Instructions:    []types.Instruction{transferInstruction},
	})

	// 8. 构造交易
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{sender}, // Provide the account (which contains private key) as signer
		Message: message,
	})
	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err)
	}

	// 9. 发送交易
	txhash, err := rpcClient.SendTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf(" Transaction sent with signature: %s\n", txhash)

}
