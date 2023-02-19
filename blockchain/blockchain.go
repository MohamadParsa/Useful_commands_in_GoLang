package main

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	chain             []*Block
	transactionPool   []*Transaction
	blockChainAddress string
}

func (b *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash, b.transactionPool)
	b.chain = append(b.chain, block)
	b.transactionPool = []*Transaction{}
	return block
}

func (b *BlockChain) AddTransaction(transaction *Transaction) {
	b.transactionPool = append(b.transactionPool, transaction)
}

func (b *BlockChain) MakeCopyTransactionPool() []*Transaction {
	transactionPool := []*Transaction{}
	for _, tr := range b.transactionPool {
		temp := *tr
		transactionPool = append(transactionPool, &temp)
	}
	return transactionPool
}

func (b *BlockChain) GetLastBlock() *Block {
	return b.chain[len(b.chain)-1]
}

func (b *BlockChain) Print() {
	b.PrintChain()
	b.PrintTransactionPool()
}

func (b *BlockChain) PrintChain() {
	fmt.Println(strings.Repeat("#", 20), " Block Chain ", strings.Repeat("#", 20))
	for i, v := range b.chain {
		fmt.Println(strings.Repeat("-", 10), "Chain ", i+1, strings.Repeat("-", 10))
		v.Print()
	}
	fmt.Println(strings.Repeat("_", 40))
}

func (b *BlockChain) PrintTransactionPool() {
	fmt.Println(strings.Repeat("=", 20), " Transactions Pool", strings.Repeat("=", 20))

	for i, v := range b.transactionPool {
		fmt.Println(strings.Repeat("$", 20), " Transaction Pool Index ", i+1, strings.Repeat("$", 20))
		fmt.Println(v)
	}
	fmt.Println(strings.Repeat("_", 40))
}

func NewBlockChain() *BlockChain {
	blockChain := &BlockChain{}
	block := Block{}
	blockChain.CreateBlock(0, block.Hash())
	return blockChain
}
