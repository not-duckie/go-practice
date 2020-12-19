package main

import ("fmt")


func reverseArray(a []int32) []int32{
	n := len(a)-1
	for i:=0;i<len(a)/2;i++{
		a[i],a[n-i] = a[n-i],a[i]
	}
	return a

}

func main(){
	fmt.Println(reverseArray([]int32{1,2,3,4,5,6,7,8,9,10}))

}
