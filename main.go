package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockchain() *Blockchain {
	blockchain := new(Blockchain)
	blockchain.BuildBlock(0, "Genesis")
	return blockchain
}

func (blockchain *Blockchain) BuildBlock(nonce int, previousHash string) *Block {
	block := NewBlock(nonce, previousHash)
	blockchain.chain = append(blockchain.chain, block)
	return block
}

func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain {
		fmt.Printf("%s Block %d  %s\n",strings.Repeat("$", 15), i, strings.Repeat("$", 15))
		block.Print()
	}
	fmt.Printf("%s", strings.Repeat("#", 50))
}

func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.timestamp = time.Now().Unix()
	block.nonce = nonce
	block.previousHash = previousHash
	return block
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func (block *Block) Print() {
	fmt.Printf("timestamp	%d\n", block.timestamp)
	fmt.Printf("nonce		%d\n", block.nonce)
	fmt.Printf("previousHash	%s\n", block.previousHash)
	fmt.Printf("transactions	%s\n", block.transactions)
}

func main() {
	blockchain := NewBlockchain()
	blockchain.BuildBlock(7, "1")
	blockchain.BuildBlock(9, "2")
	blockchain.BuildBlock(12, "3")
	blockchain.Print()
}