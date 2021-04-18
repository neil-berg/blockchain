package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/neil-berg/blockchain/util"
)

// Difficulty is a static number here, but would be dynamically increasing in
// reality to account for increased miners and computational power over time
const Difficulty = 12

// ProofOfWork is the shape of a block's proof of work
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// CreateData returns a single slice of bytes of the POW's data consisting of the
// block's previous hash and data, along with the current nonce and the difficulty.
func (pow *ProofOfWork) CreateData(nonce int) []byte {
	nonceBytes, err := util.NumToBytes(int64(nonce))
	if err != nil {
		log.Panic(err)
	}
	difficultyBytes, err := util.NumToBytes(Difficulty)
	if err != nil {
		log.Panic(err)
	}
	timestampBytes, err := util.NumToBytes(pow.Block.Timestamp.Unix())
	if err != nil {
		log.Panic(err)
	}

	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			nonceBytes,
			difficultyBytes,
			timestampBytes,
		},
		[]byte{},
	)
	return data
}

// Run loops over a nearly inifnite range of nonces, computing hashes with each
// until the hash's big int representation is less than the target hash's big
// int representation. When that happens, we declare the block to be signed.
func (pow *ProofOfWork) Run() (int, [32]byte) {
	var initHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.CreateData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		// Convert the hash to a big int
		initHash.SetBytes(hash[:])
		// Compare the big int hash to the target big int hash:
		if initHash.Cmp(pow.Target) == -1 {
			// Computed hash is less than the target hash, so we signed the block
			break
		} else {
			// Increment nonce and try again
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash
}

// Validate takes the block's nonce, recomputes the block's hash, and confirms
// that the hash as a big int is less than the POW's target.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.CreateData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// NewProof does a new proof
func NewProof(block *Block) *ProofOfWork {
	// Initialize the target at 1, then left shift it by the difficulty
	target := big.NewInt(1)
	// 256 is the number of bits in the block's hash (sha256). As difficulty
	// increases, the target value decreases, making it harder to complete the proof.
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{block, target}
	return pow
}
