package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type coinToss struct {
	sync.Mutex
	values []int
	score  int
}

func main() {
	// coinTossGameWithRace implemented with data race condition
	// if you run app with : "go run ." you don't see data race warning
	// type "go run -race ." to see data race warning

	// result with "go run -race ." :
	//		game score with data race:       6  , values:  [1 0 1 1 1 0 0 1]
	//		game score without data race:    3  , values:  [0 0 0 1 0 1 0 1 0 0]
	// result with "go run ." :
	//		game score with data race:       7  , values:  [0 1 1 1 1 1 0 0 1 1]
	//		game score without data race:    5  , values:  [1 0 1 1 0 1 0 1 0 0]

	coinTossGameWithRace()
	coinTossGameWithoutRace()

}

func coinTossGameWithRace() {
	c := coinToss{}
	var w sync.WaitGroup
	gameSets := 10
	w.Add(gameSets)
	for i := 0; i < gameSets; i++ {
		go func() {
			defer w.Done()
			randomNumber := getRandomNumber()
			c.values = append(c.values, randomNumber)
			if randomNumber == 1 {
				c.score++
			}
		}()
	}
	w.Wait()
	fmt.Println("game score with data race:	", c.score, " , values: ", c.values)
}
func coinTossGameWithoutRace() {
	c := coinToss{}
	var w sync.WaitGroup
	gameSets := 10
	w.Add(gameSets)
	for i := 0; i < gameSets; i++ {
		go func() {
			defer w.Done()
			randomNumber := getRandomNumber()
			c.Lock()
			c.values = append(c.values, randomNumber)
			if randomNumber == 1 {
				c.score++
			}
			c.Unlock()
		}()
	}
	w.Wait()
	fmt.Println("game score without data race:	", c.score, " , values: ", c.values)

}
func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}
