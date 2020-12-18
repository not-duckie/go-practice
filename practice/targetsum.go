package main

import "fmt"

func TwoSum(numbers []int,target int)(res [2]int) {
	for i := range numbers{
		for j:=len(numbers)-1; j>i;j--{
			if(numbers[i]+numbers[j]==target){
				return [...]int{i,j}
			}
		}
	}
	return [...]int{0,0}
}
func main(){
	fmt.Println(TwoSum([]int{1,2,3},4))

}