package main

import "fmt"

func main() {
	fmt.Println(solution("eye"))
	fmt.Println(solution("eye2"))

}

func solution(inputString string) bool {
	l := len(inputString)
	for i := 0; i < l/2; i++ {
		if inputString[i] != inputString[l-1-i] {
			return false
		}
	}
	return true
}
