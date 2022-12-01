package main

import (
	"fmt"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
)

type BlockChain struct {
	Chain           []*Block
	transactionPool []*Transaction
}

func (b *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash, b.transactionPool)
	b.Chain = append(b.Chain, block)
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

func (b *BlockChain) ValidProof(nonce int, nonceChan chan int) {
	zeros := strings.Repeat("0", MINING_DIFFICULTY)

	guessBlock := Block{
		nonce:        nonce,
		previousHash: b.GetLastBlock().Hash(),
		transactions: b.MakeCopyTransactionPool()}
	if zeros == guessBlock.HashString()[:MINING_DIFFICULTY] {
		nonceChan <- nonce
	}
}

func (b *BlockChain) ProofOfWork() int {
	nonceChan := make(chan int)
	nonce := 0
	for {

		select {
		case v, ok := <-nonceChan:
			if ok {
				return v
			}
		default:
			go b.ValidProof(nonce, nonceChan)
		}
		nonce++
	}

}

func (b *BlockChain) GetLastBlock() *Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *BlockChain) Print() {
	b.PrintChain()
	b.PrintTransactionPool()
}

func (b *BlockChain) PrintChain() {
	fmt.Println(strings.Repeat("#", 20), " Block Chain ", strings.Repeat("#", 20))
	for i, v := range b.Chain {
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
