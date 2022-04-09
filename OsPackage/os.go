package main

import (
	"log"
	"os"
)

func main() {
	// reaturn current directory address
	currentDirectoryAddress, err := os.Getwd()
	log.Println(currentDirectoryAddress, err)

	os.Exit(1)
}
