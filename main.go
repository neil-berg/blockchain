package main

import (
	"github.com/neil-berg/blockchain/blockchain"
)

func main() {
	// db := database.Open()
	// defer db.Close()

	// key := []byte("some-key")
	// value := []byte("this-value-for-some-key2")
	// err := db.Write(database.BlocksBucket, key, value)
	// v, err := db.Read(database.BlocksBucket, key)
	// fmt.Println("the value is......", v)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	chain := blockchain.Init()

	chain.AddBlock("First block")
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
