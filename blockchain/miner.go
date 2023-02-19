package main

import "strings"

const (
	MINING_ADDRESS    = "Miner"
	MINING_REWARD     = 1.0
	MINING_DIFFICULTY = 3
)

type Miner struct {
	blockChain *BlockChain
}

func (m *Miner) ValidProof(nonce int, nonceChan chan int) {
	zeros := strings.Repeat("0", MINING_DIFFICULTY)

	guessBlock := Block{
		nonce:        nonce,
		previousHash: m.blockChain.GetLastBlock().Hash(),
		transactions: m.blockChain.MakeCopyTransactionPool()}
	if zeros == guessBlock.HashString()[:MINING_DIFFICULTY] {
		nonceChan <- nonce
	}
}

func (m *Miner) ProofOfWork() int {
	nonceChan := make(chan int)
	nonce := 0
	for {

		select {
		case v, ok := <-nonceChan:
			if ok {
				return v
			}
		default:
			go m.ValidProof(nonce, nonceChan)
		}
		nonce++
	}

}
