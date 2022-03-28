package blockchain

import (
	"gopherchain/block"
	"gopherchain/chain"
	"gopherchain/transaction"
	"strings"
	"fmt"
)

type Blockchain struct {
	transactionPool []*transaction.Transaction
	chain chain.Chain
}

func NewBlockchain() *Blockchain {
	
	blockchain := new(Blockchain)
	blockchain.BuildFirstBlock(0)
	return blockchain
}

func (blockchain *Blockchain) BuildBlock(nonce int) *block.Block {

	lastBlock := blockchain.chain.Peek()
	block := block.NewBlock(nonce, lastBlock.Hash(), blockchain.transactionPool)
	blockchain.transactionPool = []*transaction.Transaction{}
	blockchain.chain.Add(block)
	return block
}

func (blockchain *Blockchain) BuildFirstBlock(nonce int) *block.Block {
	genesisBlock := &block.Block{}
	block := block.NewBlock(nonce, genesisBlock.Hash(), blockchain.transactionPool)
	blockchain.chain.Add(block)
	return block
}

func (blockchain *Blockchain) AddTransaction (sender string, recipient string, value float32) {
	transaction := transaction.NewTransaction(sender, recipient, value)
	blockchain.transactionPool = append(blockchain.transactionPool, transaction)
}

func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain.Blocks() {
		fmt.Printf("%s Block %d  %s\n",strings.Repeat("$", 50), i, strings.Repeat("$", 50))
		block.Print()
	}
	fmt.Printf("%s", strings.Repeat("#", 110))
}