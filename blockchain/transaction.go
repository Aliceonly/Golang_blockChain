package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Transaction struct {
	ID      string
	Inputs  []TXInput
	Outputs []TXOutput
}

type TXInput struct {
	Txid      string
	OutputIdx int
	Signature string
}

type TXOutput struct {
	Value  float64
	PubKey string
}

func NewTransaction(inputs []TXInput, outputs []TXOutput) *Transaction {
	tx := &Transaction{Inputs: inputs, Outputs: outputs}
	tx.ID = tx.calculateHash()
	return tx
}

func (tx *Transaction) calculateHash() string {
	var inputsHashes, outputsHashes string
	for _, input := range tx.Inputs {
		inputsHashes += input.Txid + input.Signature
	}
	for _, output := range tx.Outputs {
		outputsHashes += strconv.FormatFloat(output.Value, 'f', -1, 64) + output.PubKey
	}
	data := hex.EncodeToString([]byte(inputsHashes + outputsHashes))
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

func (tx *Transaction) VerifySignature() bool {
	// 这里添加您的签名验证逻辑
	// 由于实现签名涉及到密钥管理、签名算法等复杂过程，您可能需要引入额外的库
	return true
}

func (bc *Blockchain) IsValidTransaction(tx *Transaction) bool {
	if !tx.VerifySignature() {
		return false
	}
	// 您还可以添加其他验证逻辑，例如验证输入输出总额平衡等
	return true
}