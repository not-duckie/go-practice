package main

import "fmt"

func max(res,count int)int{
	if res>count{
		return res
	} else {return count}
}

func findMaxConsecutiveOnes(nums []int)(res int){
	count := 0
	for _,i := range nums{
		if(i==1){
			count++
		} else { count=0 }
		res = max(res,count)
	}

	return res
}

func main(){
	fmt.Println(findMaxConsecutiveOnes([]int{1}))
}
