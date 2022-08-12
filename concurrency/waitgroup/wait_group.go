package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := sync.WaitGroup{}

	list := []string{
		"Number -> 0",
		"Number -> 1",
		"Number -> 2",
		"Number -> 3",
		"Number -> 4",
		"Number -> 5",
		"Number -> 6",
		"Number -> 7",
		"Number -> 8",
		"Number -> 9",
		"Number -> 10",
		"Number -> 11",
	}

	waitGroup.Add(len(list))

	for _, text := range list {
		go concurrentPrint(text, &waitGroup)
	}

	waitGroup.Wait()

	waitGroup.Add(1)
	concurrentPrint("end .... ", &waitGroup)
}

func concurrentPrint(text string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(text)
}
