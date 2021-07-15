package main

import (
	"github.com/neil-berg/blockchain/blockchain"
	commandline "github.com/neil-berg/blockchain/cli"
	"github.com/neil-berg/blockchain/database"
)

func main() {
	db := database.Open()
	defer db.Close()

	chain := blockchain.Init(db)
	cli := commandline.CLI{Chain: chain}
	cli.Run()

	// chain.AddBlock("First block")
	// chain.AddBlock("Second block")
	// chain.AddBlock("Third block")

	// for i, block := range chain.Blocks {
	// 	fmt.Printf("============= BLOCK %d ===============\n", i)
	// 	fmt.Printf("Timestamp:\t %v\n", block.Timestamp)
	// 	fmt.Printf("Data:\t\t %s\n", block.Data)
	// 	fmt.Printf("Hash:\t\t %x\n", block.Hash)
	// 	fmt.Printf("Previous hash:\t %x\n", block.PrevHash)
	// 	fmt.Printf("Nonce: \t\t %d\n", block.Nonce)

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("Valid:\t\t %v\n", pow.Validate())
	// }

}
