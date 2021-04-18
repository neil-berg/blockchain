package blockchain

// Block shape
type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
	Nonce    int
}

// Blockchain shape
type Blockchain struct {
	Blocks []*Block
}

// CreateBlock creates a new block given its data and previous block's hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
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
