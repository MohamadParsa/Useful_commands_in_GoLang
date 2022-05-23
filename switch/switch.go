package main

import "fmt"

func Main() {
	i := 5
	switch i {
	case 5:
		i++
		fmt.Println(i)
		fallthrough
	case 6:
		i++
		fmt.Println(i)
	case 7:
		i++
		fmt.Println(i)
	default:
		fmt.Println(100)
	}
}
