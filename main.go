package main

import (
	// "fmt"
	"log"
	// "gopherchain/blockchain"
	"gopherchain/blockchain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := blockchain.NewBlockchain()

	blockchain.AddTransaction("X", "Y", 25.0989)

	blockchain.BuildBlock(blockchain.ProofOfWork())

	blockchain.AddTransaction("C", "O", 1.0989)
	blockchain.BuildBlock(blockchain.ProofOfWork())
	blockchain.Print()
}
