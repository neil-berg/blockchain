package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block shape
type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
}

// Blockchain shape
type Blockchain struct {
	Blocks []*Block
}

// DeriveHash adds a hash to a new block based on the previous block's hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block given its data and previous block's hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash}
	block.DeriveHash()
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
