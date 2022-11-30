package main

import "fmt"

func main() {

	blockChain := NewBlockChain()
	b := blockChain.CreateBlock(100, blockChain.GetLastBlock().Hash())
	fmt.Println(b.Hash())
	b = blockChain.CreateBlock(150, blockChain.GetLastBlock().Hash())
	fmt.Println(b.Hash())

	blockChain.PrintChain()
}
