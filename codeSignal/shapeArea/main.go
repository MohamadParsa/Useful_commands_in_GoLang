// https://app.codesignal.com/arcade/intro/level-2/yuGuHvcCaFCKk56rJ

package main

import "fmt"

func main() {
	fmt.Println(solution(2))
	fmt.Println(solution(3))

}

func solution(n int) int {
	if n == 1 {
		return 1
	}
	cellCount := 1
	courrentLineCellCount := 1
	for i := 2; i <= n; i++ {
		courrentLineCellCount = courrentLineCellCount + 2
		cellCount = cellCount + courrentLineCellCount
	}

	return cellCount + (cellCount - courrentLineCellCount)
}
