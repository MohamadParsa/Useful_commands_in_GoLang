package main

import "fmt"

func main() {
	fmt.Println(solution([]int{3, 6, -2, -5, 7, 3}))
}

func solution(inputArray []int) int {

	if len(inputArray) == 1 {
		return inputArray[0]
	}
	s := 0
	if len(inputArray) > 2 {
		s = solution(inputArray[2:])
	}
	if inputArray[0]*inputArray[1] < s {
		return s
	}
	return solution(inputArray[0:1]) * solution(inputArray[1:2])
}
