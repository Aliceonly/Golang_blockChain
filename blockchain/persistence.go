package blockchain

import (
	"encoding/gob"
	"os"
)

const blockchainFileName = "blockchain.gob"

type Persistence struct {
	filePath string
}

func NewPersistence() *Persistence {
	return &Persistence{filePath: blockchainFileName}
}

func (p *Persistence) SaveBlockchain(bc *Blockchain) error {
	file, err := os.Create(p.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(bc)
}

func (p *Persistence) LoadBlockchain() (*Blockchain, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var blockchain Blockchain
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&blockchain)
	if err != nil {
		return nil, err
	}

	return &blockchain, nil
}
