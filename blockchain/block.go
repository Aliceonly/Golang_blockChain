package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	PrevHash     string
	Hash         string
	Nonce        int64
	Transactions []Transaction
}

func NewBlock(index int64, prevHash string, transactions []Transaction, nonce int64) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		PrevHash:     prevHash,
		Hash:         "",
		Nonce:        nonce,
		Transactions: transactions,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func (b *Block) calculateHash() string {
	hashData := strconv.FormatInt(b.Index, 16) + strconv.FormatInt(b.Timestamp, 16) + b.PrevHash + strconv.FormatInt(b.Nonce, 16)
	hashBytes := sha256.Sum256([]byte(hashData))
	return hex.EncodeToString(hashBytes[:])
}
