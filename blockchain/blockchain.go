package blockchain

import (
	"bytes"
	"encoding/gob"
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

// Serialize encodes a block into a gob
func (block *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		return []byte{}, err
	}
	return result.Bytes(), nil
}

// Deserialize deserializes a slice of encoded gob bytes into a Block
func Deserialize(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return &Block{}, err
	}
	return &block, nil
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
