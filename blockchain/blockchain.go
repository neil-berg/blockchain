package blockchain

import (
	"fmt"
	"time"
)

// Block shape
type Block struct {
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
	Timestamp time.Time
}

// Blockchain shape
type Blockchain struct {
	Blocks []*Block
}

// CreateBlock performs the block's proof-of-work, populating the block with a
// hash and nonce that can validated before attaching to the chain.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash, 0, time.Now()}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	fmt.Printf("Completed proof-of-work: \n \tNonce: \t%d\n \tHash: \t%x\n", nonce, hash)
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// AddBlock adds a new block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}

// Genesis returns the first block of the blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Init initializes a new blockchain
func Init() *Blockchain {
	genesis := Genesis()
	blockchain := &Blockchain{[]*Block{genesis}}
	return blockchain
}
