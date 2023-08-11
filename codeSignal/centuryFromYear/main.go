package main

import "fmt"

func main() {
	fmt.Println(centuryFromYear(1))
	fmt.Println(centuryFromYear(100))
	fmt.Println(centuryFromYear(375))

	fmt.Println(centuryFromYear(1700))
	fmt.Println(centuryFromYear(1905))

}

func centuryFromYear(year int) int {
	c := year / 100
	if year%100 != 0 {
		c++
	}
	return c
}
