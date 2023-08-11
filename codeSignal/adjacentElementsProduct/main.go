package main

import "fmt"

func main() {
	fmt.Println(solution([]int{-23, 4, -3, 8, -12}))
}

func solution(inputArray []int) int {

	if len(inputArray) == 1 {
		return inputArray[0]
	}
	s := 0
	if len(inputArray) > 2 {
		s = solution(inputArray[1:])
		if inputArray[0]*inputArray[1] < s {
			return s
		}
	}
	return inputArray[0] * inputArray[1]
}
