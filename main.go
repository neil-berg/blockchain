package main

import (
	"fmt"

	"github.com/neil-berg/blockchain/blockchain"
)

func main() {
	chain := blockchain.Init()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Block data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
	}
}
