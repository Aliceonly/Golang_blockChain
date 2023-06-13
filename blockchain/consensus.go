package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBits)
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) Run() (int64, string) {
	var nonce int64
	var hash string
	for nonce = 0; nonce < math.MaxInt64; nonce++ {
		hash = pow.calculateHashWithNonce(nonce)
		hashInt := new(big.Int).SetBytes([]byte(hash))
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
	}
	return nonce, hash
}

func (pow *ProofOfWork) calculateHashWithNonce(nonce int64) string {
	data := fmt.Sprintf("%x%x%x%x",
		pow.block.Index,
		pow.block.Timestamp,
		pow.block.PrevHash,
		nonce)
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

func (pow *ProofOfWork) Validate() bool {
	hash := pow.calculateHashWithNonce(pow.block.Nonce)
	hashInt := new(big.Int).SetBytes([]byte(hash))
	return hashInt.Cmp(pow.target) == -1
}
