package main

import (
	"fmt"
	"sort"
)

func sortedSquares(num []int) []int {
	for i := range num {
		num[i] *= num[i]
	}
	sort.Ints(num)
	return num
}

func main() {
	fmt.Println(sortedSquares([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
