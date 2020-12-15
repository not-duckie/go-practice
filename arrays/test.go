package main

import (
	"fmt"
	"strings" )

func printfunc(ref [][]int){
	ref[0][0] = 0
	fmt.Println("Printing from func")
	fmt.Println(ref)
}

func main(){
	array := []int{1,2,3}
	array := [3]int {1,2,3}
	fmt.Println(array," => array")

	//passing as copy
	copy_array := array
	for i := 0; i<3; i++{
		fmt.Println(array[i])
	}

	copy_array[2] = 69

	fmt.Println(strings.Repeat("*",20))

	for i := 0; i<3; i++{
		fmt.Println(copy_array[i])
	}


	fmt.Println(strings.Repeat("*",20))
	//passing as an reference
	ref := &array
	ref[1] = 69
	for i:=0;i<len(array);i++ {
		fmt.Println(ref[i])
	}

	fmt.Println(strings.Repeat("*",20))

	for i := 0; i<3; i++{
		fmt.Println(array[i])
	}
	array2 := [][]int{[]int{1,0,0},[]int{0,1,0},[]int{0,0,1}}
	printfunc(array2)

	fmt.Printf("%T ",array2)
}
