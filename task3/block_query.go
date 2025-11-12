package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

func main() {
	// 1. 创建 RPC 客户端连接到 Devnet
	rpcClient := client.NewClient(rpc.DevnetRPCEndpoint)

	// 2. 获取最新区块哈希
	blockhashResp, err := rpcClient.GetLatestBlockhash(context.TODO())
	if err != nil {
		log.Fatalf("Failed to get latest blockhash: %v", err)
	}
	fmt.Printf("Latest blockhash: %s\n", blockhashResp.Blockhash)

	// 3. 查询账户余额
	account := types.NewAccount()
	publicKey := account.PublicKey.ToBase58()
	fmt.Printf("Account public key: %s\n", publicKey)

	balance, err := rpcClient.GetBalance(context.TODO(), publicKey)
	if err != nil {
		log.Fatalf("Failed to get balance for %s: %v", account, err)
	}
	fmt.Printf("Balance for %s: %d Lamports (%.10f SOL)\n", account, balance, float64(balance)/float64(1e9)) // 1 SOL = 1,000,000,000 Lamports

	// 4. 获取账户详细信息
	accountInfo, err := rpcClient.GetAccountInfo(context.TODO(), publicKey)
	if err != nil {
		log.Fatalf("Failed to get account info for %s: %v", account, err)
	}
	fmt.Printf("Account Info - Lamports: %d, Owner: %s, Executable: %t, RentEpoch: %d\n",
		accountInfo.Lamports, accountInfo.Owner.String(), accountInfo.Executable, accountInfo.RentEpoch)
}
