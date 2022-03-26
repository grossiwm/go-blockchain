package blockchain

import (
	block "gopherchain/block"
	"strings"
	"fmt"
)

type Blockchain struct {
	transactionPool []string
	chain []*block.Block
}

func NewBlockchain() *Blockchain {
	blockchain := new(Blockchain)
	blockchain.BuildBlock(0, "Genesis")
	return blockchain
}

func (blockchain *Blockchain) BuildBlock(nonce int, previousHash string) *block.Block {
	block := block.NewBlock(nonce, previousHash)
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