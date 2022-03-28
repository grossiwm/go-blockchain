package blockchain

import (
	"gopherchain/block"
	"gopherchain/transaction"
	"gopherchain/chain"
	"strings"
	"fmt"
)

const MINING_DIFFICULTY = 3

type Blockchain struct {
	transactionPool []*transaction.Transaction
	chain chain.Chain
}

func NewBlockchain() *Blockchain {
	
	blockchain := new(Blockchain)
	blockchain.BuildFirstBlock(0)
	return blockchain
}

func (blockchain *Blockchain) LastBlock() *block.Block {
	lastBlock := blockchain.chain.Peek()
	return &lastBlock
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

func (blockchain *Blockchain) CopyTransactionPool() []*transaction.Transaction {
	transactions := make([]*transaction.Transaction, 0)
	for _, t := range blockchain.transactionPool {
		transactions = append(transactions, 
			transaction.NewTransaction(t.GetSenderAddress(),
			 t.GetRecipientAddress(), 
			 t.GetValue()))
	}
	return transactions
}

func (blockchain Blockchain) ValidProof(nonce int,
	 previousHash [32]byte,
	 transactions []*transaction.Transaction,
	  difficulty int ) bool {
		zeros := strings.Repeat("0", difficulty)
		guessBlock := block.NewBlock(nonce, previousHash, transactions)
		guessHash := fmt.Sprintf("%x", guessBlock.Hash())
		return guessHash[:difficulty] == zeros
}

func (blockchain *Blockchain) ProofOfWork() int {
	transactions := blockchain.CopyTransactionPool()
	lastBlock := blockchain.chain.Peek()
	previousHash := lastBlock.Hash()
	nonce := 0
	for !blockchain.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain.Blocks() {
		fmt.Printf("%s Block %d  %s\n",strings.Repeat("$", 50), i, strings.Repeat("$", 50))
		block.Print()
	}
	fmt.Printf("%s", strings.Repeat("#", 110))
}