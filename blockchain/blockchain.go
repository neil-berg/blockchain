package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/neil-berg/blockchain/database"
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
	// The blockchain's tip is the last block hash stored in the DB
	tip []byte
	// Instance of our DB
	db *database.Database
}

// Iterator shape
type Iterator struct {
	currentHash []byte
	db          *database.Database
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
func (chain *Blockchain) AddBlock(data string) error {
	tipKey := []byte(database.TipKey)

	tip, err := chain.db.Read(database.BlocksBucket, tipKey)
	if err != nil {
		log.Fatal(err)
	}

	block := CreateBlock(data, tip)
	serializedBlock, err := block.Serialize()
	err = chain.db.Write(database.BlocksBucket, block.Hash, serializedBlock)
	err = chain.db.Write(database.BlocksBucket, tipKey, block.Hash)
	if err != nil {
		return err
	}

	fmt.Println("============= ADDED BLOCK ===============")
	fmt.Printf("Timestamp:\t %v\n", block.Timestamp)
	fmt.Printf("Data:\t\t %s\n", block.Data)
	fmt.Printf("Hash:\t\t %x\n", block.Hash)
	fmt.Printf("Previous hash:\t %x\n", block.PrevHash)
	fmt.Printf("Nonce: \t\t %d\n", block.Nonce)
	return nil
}

// GetNewIterator returns a new blockchain iterator
func (chain *Blockchain) GetNewIterator() *Iterator {
	return &Iterator{chain.tip, chain.db}
}

// Next returns the next block in a blockchain iterator
func (iterator *Iterator) Next() *Block {
	var block *Block

	data, err := iterator.db.Read(database.BlocksBucket, iterator.currentHash)
	if err != nil {
		fmt.Println("Failed to get next block")
	}

	block, err = Deserialize(data)
	if err != nil {
		fmt.Println("Failed to deserialize the block")
	}

	iterator.currentHash = block.PrevHash

	return block
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
func Init(db *database.Database) *Blockchain {
	var tip []byte
	tipKey := []byte(database.TipKey)

	emtpyBlocksBucket := db.EmptyBucket(database.BlocksBucket)

	if emtpyBlocksBucket {
		// Create and store the genesis block and it's hash as the new chain's tip
		genesis := Genesis()
		fmt.Println("creating genesis block")
		serializedBlock, err := genesis.Serialize()
		err = db.Write(database.BlocksBucket, genesis.Hash, serializedBlock)
		fmt.Println("storing genisis")
		err = db.Write(database.BlocksBucket, tipKey, genesis.Hash)
		fmt.Println("storing tip")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("reading tip")
		// Blockchain exists, just read the tip
		lastHash, err := db.Read(database.BlocksBucket, tipKey)
		if err != nil {
			log.Fatal(err)
		}
		tip = lastHash
		fmt.Printf("Tip: %x\n", tip)
	}

	return &Blockchain{tip, db}
}
