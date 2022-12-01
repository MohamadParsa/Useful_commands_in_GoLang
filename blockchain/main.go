package main

import "fmt"

func main() {
	tr := NewTransaction("sende", "recipient", "hi")

	blockChain := NewBlockChain()
	blockChain.AddTransaction(tr)

	nonce := blockChain.ProofOfWork()
	fmt.Println(nonce)
	blockChain.CreateBlock(nonce, blockChain.GetLastBlock().Hash())
	nonce = blockChain.ProofOfWork()

	blockChain.CreateBlock(nonce, blockChain.GetLastBlock().Hash())

	blockChain.Print()

}
