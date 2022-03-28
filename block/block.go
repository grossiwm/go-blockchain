package block

import (
	"fmt"
	"time"
	"encoding/json"
	"crypto/sha256"
	"gopherchain/transaction"
)

type Block struct {
	nonce int
	previousHash [32]byte
	timestamp int64
	transactions []*transaction.Transaction
}

func (block *Block) MarshalJSON() ([]byte, error) {
	type B struct {
		Timestamp int64 `json:"timestamp"`
		Nonce int `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Transactions []*transaction.Transaction `json:"transactions"`
	}

	return json.Marshal(&B {
		Timestamp: block.timestamp,
		Nonce: block.nonce,
		PreviousHash: block.previousHash,
		Transactions: block.transactions,
	})
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
	block := new(Block)
	block.timestamp = time.Now().Unix()
	block.nonce = nonce
	block.previousHash = previousHash
	block.transactions = transactions
	return block
}

func (block *Block) Print() {
	fmt.Printf("timestamp	%d\n", block.timestamp)
	fmt.Printf("nonce		%d\n", block.nonce)
	fmt.Printf("previousHash	%x\n", block.previousHash)
	for _, t := range block.transactions {
		t.Print()
	}
}

func (block *Block) Hash() [32]byte {
	marshalled, _ := json.Marshal(block)
	return sha256.Sum256([]byte(marshalled));
	
}