package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.timestamp = time.Now().Unix()
	block.nonce = nonce
	block.previousHash = previousHash
	return block
}

func (block *Block) Print() {
	fmt.Printf("timestamp	%d\n", block.timestamp)
	fmt.Printf("nonce		%d\n", block.nonce)
	fmt.Printf("previousHash	%s\n", block.previousHash)
	fmt.Printf("transactions	%s\n", block.transactions)
}