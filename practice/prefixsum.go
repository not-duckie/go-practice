package main

import "fmt"

func calculation(num int)([]int){
	arr := []int{1,2,3,4,5,6,7,8,9}
	sum := make([]int,9)
	sum[0] = arr[0]
	for i:=1;i<num+1;i++{
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

func sumtill(n int){
	sum := make([]int,9)
	if sum[n]==0{
		sum = calculation(n)
		fmt.Println(sum[n])
	} else { fmt.Println(sum[n]) }
}

func main(){
	sumtill(3)
	sumtill(8)
	sumtill(3)
	sumtill(2)
	sumtill(3)
	sumtill(8)
	sumtill(3)
	sumtill(2)
	sumtill(3)
	sumtill(3)
	sumtill(3)
	sumtill(8)
	sumtill(3)
	sumtill(2)
	sumtill(8)
	sumtill(3)
	sumtill(2)
	sumtill(8)
	sumtill(3)
	sumtill(2)
}
