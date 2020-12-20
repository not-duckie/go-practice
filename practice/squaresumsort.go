package main

import (
	"fmt"
	"sort"
)

func sortedSquares(num []float32) []float32 {
	for i := range num {
		num[i] *= num[i]
	}
	sort.Float32s(num)
	return num
}

func main() {
	fmt.Println(sortedSquares([]float32{10, 9.8, 8, 7, 6, 5, 4, 3, 2, 1}))
}
