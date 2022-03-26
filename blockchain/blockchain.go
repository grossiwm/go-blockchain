package blockchain

import (
	"gopherchain/block"
	"gopherchain/chain"
	"strings"
	"fmt"
)

type Blockchain struct {
	transactionPool []string
	chain chain.Chain
}

func NewBlockchain() *Blockchain {
	
	blockchain := new(Blockchain)
	blockchain.BuildFirstBlock(0)
	return blockchain
}

func (blockchain *Blockchain) BuildBlock(nonce int) *block.Block {

	lastBlock := blockchain.chain.Peek()
	block := block.NewBlock(nonce, lastBlock.Hash())
	blockchain.chain.Add(block)
	return block
}

func (blockchain *Blockchain) BuildFirstBlock(nonce int) *block.Block {
	genesisBlock := &block.Block{}
	block := block.NewBlock(nonce, genesisBlock.Hash())
	blockchain.chain.Add(block)
	return block
}

func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain.Blocks() {
		fmt.Printf("%s Block %d  %s\n",strings.Repeat("$", 50), i, strings.Repeat("$", 50))
		block.Print()
	}
	fmt.Printf("%s", strings.Repeat("#", 110))
}