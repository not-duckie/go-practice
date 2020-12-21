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

func caching(n int){
	sum := make([]int,9)
	if sum[n]==0{
		sum = calculation(n)
		fmt.Println(sum[n])
	} else { fmt.Println(sum[n]) }
}

func main(){
	caching(3)
	caching(8)
	caching(3)
	caching(2)
	caching(3)
	caching(8)
	caching(3)
	caching(2)
	caching(3)
	caching(3)
	caching(3)
	caching(8)
	caching(3)
	caching(2)
	caching(8)
	caching(3)
	caching(2)
	caching(8)
	caching(3)
	caching(2)
}
