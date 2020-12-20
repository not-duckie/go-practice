package main

import "fmt"

func check(i int)bool{
	tmp:=0
	for i!=0{
		i = i/10
		tmp++ }
	if tmp%2==0{ return true}
	return false
}

func findNumbers(nums []int)int{
	count :=0
	for _,i:=range nums{
		if a:=check(i);a{
			count++
		}
	}
	return count
}

func main(){
	fmt.Println(findNumbers([]int{12,345,2,6,7896}))
}
