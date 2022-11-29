package main

import (
	"fmt"
	"time"
)

func main() {

	s := make(chan int)
	m := make(chan int)

	go secondChecker(s)
	go minuteChecker(m)

	printService(s, m)
}

func secondChecker(c chan int) {
	fmt.Println("secondChecker called")
	for {
		second := time.Now().Second()
		if second%5 == 0 {
			c <- second
			time.Sleep(1 * time.Second)
		}
	}
}
func minuteChecker(c chan int) {
	fmt.Println("minuteChecker called")
	for {
		minute := time.Now().Minute()
		if minute%2 == 0 {
			c <- minute
			time.Sleep(1 * time.Minute)
		}
	}
}
func printService(s, m chan int) {
	fmt.Println("printService called")
	for {
		select {
		case v, ok := <-s:
			if ok {
				fmt.Println("second receives : ", v)
			}
		case v, ok := <-m:
			if ok {
				fmt.Println("minute received : ", v)
			}

		}
	}
}
