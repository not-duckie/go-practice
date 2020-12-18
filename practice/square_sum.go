package main

import ("fmt")

func SquareSum(nums []int)(sum int) {
	for _,i := range nums {
		sum +=i*i
	}
	return sum
}


func main(){
	fmt.Println(SquareSum([]int{1,2,2}))

}
