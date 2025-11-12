// events.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/blocto/solana-go-sdk/types"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	ctx := context.Background()

	// 连接到 Devnet 的 WebSocket
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		log.Fatalf("Failed to connect to websocket: %v", err)
	}
	defer wsClient.Close()

	// 订阅特定账户的变更 ---
	account := types.NewAccount()
	publicKey, err := solana.PublicKeyFromBase58(account.PublicKey.ToBase58())
	if err != nil {
		log.Fatalf("Failed to get publish key: %v", err)
	}

	sub, err := wsClient.AccountSubscribe(
		publicKey,
		rpc.CommitmentRecent,
	)
	if err != nil {
		log.Fatalf("Failed to subscribe to account: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Printf("Subscribed to account %s using gagliardetto/ws client\n", account)
	fmt.Println("Listening for account updates... Press Ctrl+C to exit.")
	for {
		got, err := sub.Recv(ctx)
		if err != nil {
			log.Fatalf("Error receiving subscription: %v", err)
		}
		fmt.Println("--- Account Update Received ---")
		fmt.Printf("Account %s updated at slot %d\n", got.Value, got.Context.Slot)
		time.Sleep(100 * time.Millisecond) // 防止循环过快
	}
}
