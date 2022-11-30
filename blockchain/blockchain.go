package main

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	Chain           []*Block
	transactionPool []string
}

func (b *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash)
	b.Chain = append(b.Chain, block)
	return block
}

func (b *BlockChain) GetLastBlock() *Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *BlockChain) PrintChain() {
	fmt.Println(strings.Repeat("#", 40))
	for i, v := range b.Chain {
		fmt.Println(strings.Repeat("-", 10), "Chain ", i, strings.Repeat("-", 10))
		v.Print()
	}
	fmt.Println(strings.Repeat("#", 40))
}
func (b *BlockChain) PrintTransactionPool() {
	fmt.Println(strings.Repeat("#", 40))

	for i, v := range b.transactionPool {
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("Transaction Index:		", i)
		fmt.Println(v)
	}
	fmt.Println(strings.Repeat("#", 40))
}

func NewBlockChain() *BlockChain {
	blockChain := &BlockChain{}
	block := Block{}
	blockChain.CreateBlock(0, block.Hash())
	return blockChain
}
