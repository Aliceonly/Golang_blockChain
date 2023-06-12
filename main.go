package main

import (
	"cjjBlockChain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("Hello World")
	bc.AddBlock("Goodbye World")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
