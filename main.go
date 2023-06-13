package main

// func main() {
// 	persistence := blockchain.NewPersistence()
// 	bc, err := persistence.LoadBlockchain()
// 	if err != nil {
// 		fmt.Println("未能从文件中加载区块链，正在创建新的区块链")
// 		bc = blockchain.NewGenesisBlock()
// 	}

// 	// 启动P2P网络并同步区块链

// 	err = persistence.SaveBlockchain(bc)
// 	if err != nil {
// 		log.Panic("无法将区块链保存到文件")
// 	}
// }

import (
	"cjjBlockChainSys/blockchain" // 更改为您的项目命名空间
	"fmt"
	"log"
)

func main() {
	persistence := blockchain.NewPersistence()
	bc, err := persistence.LoadBlockchain()
	if err != nil {
		fmt.Println("未能从文件中加载区块链，正在创建新的区块链")
		bc = blockchain.NewBlockchain()
	}

	defer func() {
		err = persistence.SaveBlockchain(bc)
		if err != nil {
			log.Panic("无法将区块链保存到文件")
		}
	}()

	cli := blockchain.NewCommandLine(bc)
	cli.Run()
}
