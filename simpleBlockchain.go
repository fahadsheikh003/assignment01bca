package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
)

type Block struct {
	PreviousHash string
	Previous     *Block
	data         string
	nonce        int
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := new(Block)
	block.data = transaction
	block.Previous = nil
	block.nonce = nonce
	block.PreviousHash = previousHash
	return block
}

func GetNonce() int {
	return rand.Intn(100000000)
}

func (block *Block) ListBlocks() {
	iter := block
	for iter != nil {
		fmt.Printf("%s\n", strings.Repeat("=", 80))
		fmt.Printf("Data: %s\n", iter.data)
		fmt.Printf("Hash: %s\n", iter.CalculateHash())
		fmt.Printf("Nonce: %d\n", iter.nonce)
		if iter.PreviousHash != "" {
			fmt.Printf("Previous Hash: %s\n", iter.PreviousHash)
		}
		iter = iter.Previous
	}
}

func (block *Block) ChangeBlock(transaction string) {
	block.data = transaction
}

func (block *Block) VerifyChain(lastBlockHash string) bool {
	hash := lastBlockHash
	iter := block
	for iter != nil {
		if hash != iter.CalculateHash() {
			return false
		}
		hash = iter.PreviousHash
		iter = iter.Previous
	}
	return true
}

func (block *Block) CalculateHash() string {
	return CalculateHash(fmt.Sprintf("%d", block.nonce) + block.data + block.PreviousHash)
}
