package main

import("fmt")


func binary(low int, high int, x int, ar [10]int) int {
	if (high < low){
		return -1
	}
	mid := (low+high)/2
	if (ar[mid] < x){
		return binary(mid+1,high,x,ar)
	} else if (ar[mid]>x){
		return binary(low,mid-1,x,ar)
	} else {
		fmt.Println("Found the element at ",mid+1)
		return 1
	}
}


func main(){
	ar := [...]int{1,2,3,4,5,6,7,8,9,10}
	x := 3
	binary(0,9,x,ar)
}
