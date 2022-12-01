package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	previousHash [32]byte
	timestamp    int64
	nonce        int
	transactions []*Transaction
}

func (b *Block) Print() {
	fmt.Printf("timestamp:		%d\n", b.timestamp)
	fmt.Printf("nonce:			%d\n", b.nonce)
	fmt.Printf("previous hash:		%x\n", b.previousHash)

	for i, v := range b.transactions {
		fmt.Println(strings.Repeat("$", 20), " Transaction Index: ", i+1, strings.Repeat("$", 20))
		fmt.Println(v)
	}
}

func (b *Block) Hash() [32]byte {
	jsonByte, err := json.Marshal(b)
	if err != nil {
		log.Printf("error on marshaling")
		return [32]byte{}
	}
	return sha256.Sum256(jsonByte)
}

func (b *Block) HashString() string {
	hash := b.Hash()
	return fmt.Sprintf("%x", hash)
}
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		Transactions []*Transaction `json:"transactions"`
	}{
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		Transactions: b.transactions,
	})
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		transactions: transactions,
	}
}
