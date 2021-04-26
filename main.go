package main

import (
	"fmt"
	"log"

	"github.com/neil-berg/blockchain/blockchain"
	"github.com/neil-berg/blockchain/database"
)

func main() {
	db := database.Open()
	defer db.Close()

	key := []byte("foo")
	value := []byte("barwww")
	err := database.Write(db, key, value)
	v, err := database.Read(db, key)
	fmt.Println("the value is......", v)
	if err != nil {
		log.Fatal(err)
	}

	return
	chain := blockchain.Init()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for i, block := range chain.Blocks {
		fmt.Printf("============= BLOCK %d ===============\n", i)
		fmt.Printf("Timestamp:\t %v\n", block.Timestamp)
		fmt.Printf("Data:\t\t %s\n", block.Data)
		fmt.Printf("Hash:\t\t %x\n", block.Hash)
		fmt.Printf("Previous hash:\t %x\n", block.PrevHash)
		fmt.Printf("Nonce: \t\t %d\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("Valid:\t\t %v\n", pow.Validate())
	}
}
