package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployAndInteract() {
	// 配置
	infuraURL := "https://sepolia.infura.io/v3/a8bd9406b3ba410f8033de441bbb5420"
	privateKeyHex := "a594910ec84a2ea0dc9117ae2f2f16281a7717893730b3b97e4dfc436ad3d058" // 替换为您的私钥

	// 连接到 Sepolia
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal("连接到以太坊网络失败:", err)
	}
	defer client.Close()

	// 创建交易授权
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("私钥格式错误:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法提取公钥")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("部署者地址: %s\n", fromAddress.Hex())

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取nonce失败:", err)
	}

	// 获取 gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取gas价格失败:", err)
	}

	// 创建部署授权
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 3000000
	auth.GasPrice = gasPrice

	// 部署合约
	fmt.Println("正在部署合约...")
	address, tx, contractInstance, err := DeployCounter(auth, client)
	if err != nil {
		log.Fatal("部署合约失败:", err)
	}

	fmt.Printf("合约部署交易已发送，交易哈希: %s\n", tx.Hash().Hex())
	fmt.Printf("合约地址: %s\n", address.Hex())

	// 等待交易确认
	fmt.Println("等待交易确认...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal("等待交易确认失败:", err)
	}

	if receipt.Status != 1 {
		log.Fatal("交易执行失败")
	}

	fmt.Printf("合约部署成功！确认块号: %d\n", receipt.BlockNumber)

	// 与合约交互
	fmt.Println("\n=== 开始合约交互 ===")

	// 1. 获取初始计数器值
	count, err := contractInstance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Printf("获取计数器值失败: %v", err)
	} else {
		fmt.Printf("初始计数器值: %s\n", count.String())
	}

	// 2. 增加计数器
	fmt.Println("增加计数器...")
	incTx, err := contractInstance.Increment(auth)
	if err != nil {
		log.Printf("调用Increment失败: %v", err)
	} else {
		fmt.Printf("增加操作交易哈希: %s\n", incTx.Hash().Hex())

		// 等待交易确认
		ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel2()

		receipt2, err := bind.WaitMined(ctx2, client, incTx)
		if err != nil {
			log.Printf("等待增加操作确认失败: %v", err)
		} else if receipt2.Status == 1 {
			fmt.Println("增加操作成功")
		}
	}

	// 3. 再次获取计数器值
	count, err = contractInstance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Printf("获取计数器值失败: %v", err)
	} else {
		fmt.Printf("增加后计数器值: %s\n", count.String())
	}

	// 4. 再增加一次
	fmt.Println("再次增加计数器...")
	incTx2, err := contractInstance.Increment(auth)
	if err != nil {
		log.Printf("调用Increment失败: %v", err)
	} else {
		fmt.Printf("增加操作交易哈希: %s\n", incTx2.Hash().Hex())

		ctx3, cancel3 := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel3()

		receipt3, err := bind.WaitMined(ctx3, client, incTx2)
		if err != nil {
			log.Printf("等待增加操作确认失败: %v", err)
		} else if receipt3.Status == 1 {
			fmt.Println("增加操作成功")
		}
	}

	// 5. 获取最终计数器值
	count, err = contractInstance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Printf("获取计数器值失败: %v", err)
	} else {
		fmt.Printf("最终计数器值: %s\n", count.String())
	}

	// 6. 获取减少后的计数器值
	count, err = contractInstance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Printf("获取计数器值失败: %v", err)
	} else {
		fmt.Printf("减少后计数器值: %s\n", count.String())
	}

	fmt.Println("\n合约交互演示完成！")
}

func main() {
	deployAndInteract()
}
