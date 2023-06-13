package blockchain

import (
	"encoding/gob"
	"errors"
	"fmt"
)

func init() {
	gob.Register(&Block{})
	gob.Register(&Transaction{})
	// 将此行添加到其他结构体注册之间
}

const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := createGenesisBlock()
	bc := &Blockchain{Blocks: []*Block{genesisBlock}}
	return bc
}

func createGenesisBlock() *Block {
	coinbase := NewCoinbaseTransaction("GenesisPublicKey")
	return NewBlock(0, "", []Transaction{*coinbase}, 0)
}

func NewCoinbaseTransaction(address string) *Transaction {
	txin := TXInput{Txid: " ", OutputIdx: -1, Signature: "nil"}
	txout := TXOutput{Value: 10000, PubKey: string(address)}
	tx0 := Transaction{ID: "nil", Inputs: []TXInput{txin}, Outputs: []TXOutput{txout}}
	return &tx0
}

func (bc *Blockchain) AddBlock(newBlock *Block) error {
	// 验证新区块
	if err := bc.IsValidNewBlock(newBlock); err != nil {
		return err
	}

	bc.Blocks = append(bc.Blocks, newBlock)
	return nil
}

func (bc *Blockchain) IsValidNewBlock(newBlock *Block) error {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]

	// 验证索引
	if newBlock.Index != previousBlock.Index+1 {
		return errors.New("Invalid block index")
	}

	// 验证前一个哈希
	if newBlock.PrevHash != previousBlock.Hash {
		return errors.New("Invalid previous block hash")
	}

	// 验证新哈希
	if newBlock.Hash != newBlock.calculateHash() {
		return errors.New("Invalid block hash")
	}

	// 验证交易
	for i, tx := range newBlock.Transactions {
		if !bc.IsValidTransaction(&tx) {
			return fmt.Errorf("Invalid transaction at index %d", i)
		}
	}

	return nil
}

// func (bc *Blockchain) IsValidTransaction(tx *Transaction) bool {
// 	// 这里您可以添加验证交易的逻辑，稍后我们会更详细地讨论这部分
// 	return true
// }

func (bc *Blockchain) GetChainInfo() (int64, []*Block) {
	return int64(len(bc.Blocks)), bc.Blocks
}

func (bc *Blockchain) GetBlock(index int64) (*Block, error) {
	if index < 0 || index >= int64(len(bc.Blocks)) {
		return nil, errors.New("Block index out of range")
	}

	return bc.Blocks[index], nil
}