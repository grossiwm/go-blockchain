package main

import (
	"log"
	"gopherchain/blockchain"	
)



func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := blockchain.NewBlockchain()
	blockchain.BuildBlock(7, "1")
	blockchain.BuildBlock(9, "2")
	blockchain.BuildBlock(12, "3")
	blockchain.Print()
}