package main

import "fmt"

func test(a []int){
	a[0] = 69
	fmt.Println(a)
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	test(a)
	fmt.Println(a)
}
