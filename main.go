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
	blockchain.BuildFirstBlock(1)
	blockchain.AddTransaction("X", "Y", 25.0989)
	blockchain.AddTransaction("C", "O", 1.0989)
	blockchain.BuildBlock(5)
	blockchain.Print()
}
