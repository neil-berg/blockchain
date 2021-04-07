package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// Difficulty is a static number here, but would be dynamically increasing in
// reality to account for increased miners and computational power over time
const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 256 is the number of bytes in the block's hash (sha256)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
