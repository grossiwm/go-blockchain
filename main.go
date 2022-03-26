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
	blockchain.BuildBlock(5)
	blockchain.BuildBlock(3)
	blockchain.BuildBlock(1)
	blockchain.Print()
}
