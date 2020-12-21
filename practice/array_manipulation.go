package main

import "fmt"

func arrayManipulation(n int32,queries [][]int32) int64{
	arr := make([]int64,n)
	var max int64
	for i :=0;i<len(queries);i++{
			for j:=queries[i][0]-1;j<queries[i][1];j++{
				arr[j] += int64(queries[i][2])
				if arr[j] > max{
					max = arr[j]
				}
			}
	}
	return max
}

func main(){
	fmt.Println(arrayManipulation(5,[][]int32{{1,2,100},{2,5,100},{3,4,100}}))

}
