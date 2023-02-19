package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	swap(list)
	fmt.Println(list)
}
func swap(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
