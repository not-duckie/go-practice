package main

import (
	"fmt"
	"os"
	"strings"
)

func matrixmult(mat1 [5][5]int, mat2 [5][5]int, m1, n1, n2 int) {
	var res [5][5]int
	for i := 0; i < m1; i++ {
		for j := 0; j < n2; j++ {
			for k := 0; k < n1; k++ {
				res[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	printmatrix(res, m1, n2)
}

func printmatrix(mat [5][5]int, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}

}

func inputmatrix(m, n int) [5][5]int {
	var mat [5][5]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &mat[i][j])
		}
	}
	return mat
}

func main() {
	var m1, n1, m2, n2 int

	fmt.Printf("Enter the m1 >")
	fmt.Scanf("%d\n", &m1)

	fmt.Printf("Enter the n1 >")
	fmt.Scanf("%d\n", &n1)

	fmt.Printf("Enter the m2 >")
	fmt.Scanf("%d\n", &m2)

	fmt.Printf("Enter the n2 >")
	fmt.Scanf("%d\n", &n2)

	if m1 != n2 {
		fmt.Println("Not possible")
		os.Exit(0)
	}

	fmt.Println("Enter the matrix1")
	mat1 := inputmatrix(m1, n1)

	fmt.Println("Enter the matrix2")
	mat2 := inputmatrix(m2, n2)

	fmt.Println(strings.Repeat("*", 25))
	printmatrix(mat1, m1, n1)

	fmt.Println(strings.Repeat("*", 25))
	printmatrix(mat2, m2, n2)

	fmt.Println(strings.Repeat("*", 25))
	matrixmult(mat1, mat2, m1, n1, n2)
}
