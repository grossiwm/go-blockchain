package block

import (
	"fmt"
	"time"
	"encoding/json"
	"crypto/sha256"
)

type Block struct {
	nonce int
	previousHash [32]byte
	timestamp int64
	transactions []string
}

func (block *Block) MarshalJSON() ([]byte, error) {
	type B struct {
		Timestamp int64 `json:"timestamp"`
		Nonce int `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Transactions []string `json:"transactions"`
	}

	return json.Marshal(&B {
		Timestamp: block.timestamp,
		Nonce: block.nonce,
		PreviousHash: block.previousHash,
		Transactions: block.transactions,
	})
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	block := new(Block)
	block.timestamp = time.Now().Unix()
	block.nonce = nonce
	block.previousHash = previousHash
	return block
}

func (block *Block) Print() {
	fmt.Printf("timestamp	%d\n", block.timestamp)
	fmt.Printf("nonce		%d\n", block.nonce)
	fmt.Printf("previousHash	%x\n", block.previousHash)
	fmt.Printf("transactions	%s\n", block.transactions)
}

func (block *Block) Hash() [32]byte {
	marshalled, _ := json.Marshal(block)
	return sha256.Sum256([]byte(marshalled));
	
}