package blockchain

import (
	"bytes"
	"fmt"
	"time"
)

// Blockchain represents a blockchain data structure.
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain creates a new blockchain with the genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{})
	return &Blockchain{[]*Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsBlockValid(block, prevBlock *Block) bool {
	if block == nil || prevBlock == nil {
		return false
	}
	if !bytes.Equal(prevBlock.Hash, block.PrevHash) {
		return false
	}
	if !bytes.Equal(NewBlock(string(block.Data), []byte{}).Hash, block.Hash) {
		return false
	}
	return true
}

// Print prints the contents of the blockchain.
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %v\n", time.Unix(block.Timestamp, 0))
		fmt.Println()
	}
}
