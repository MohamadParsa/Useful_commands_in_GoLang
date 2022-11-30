package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Block struct {
	previousHash [32]byte
	timestamp    int64
	nonce        int
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("timestamp:		%d\n", b.timestamp)
	fmt.Printf("nonce:			%d\n", b.nonce)
	fmt.Printf("previous hash:		%x\n", b.previousHash)
	fmt.Printf("transactions:		%S\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	jsonByte, err := json.Marshal(b)
	if err != nil {
		log.Printf("error on marshaling")
		return [32]byte{}
	}
	return sha256.Sum256(jsonByte)
}
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PreviousHash [32]byte `json:"previous_hash"`
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		Transactions []string `json:"transactions"`
	}{
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		Transactions: b.transactions,
	})
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
	}
}
