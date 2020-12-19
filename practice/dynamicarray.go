package main

import "fmt"

func dynamicArray(n int32,queries [][]int32)(res []int32) {
	seqList := make([][]int32,n)
	var lastanswer int32
	for _,i := range queries{
		if i[0] == 1{
			index := ((i[1]^lastanswer)%n)
			seqList[index] = append(seqList[index],i[2])
		} else { index := ((i[1]^lastanswer)%n)
			lastanswer = seqList[index][i[2]%int32(len(seqList[index]))]
			res = append(res,lastanswer)
		}
	}
	fmt.Println(seqList)
	return res
}

func main(){
	fmt.Println(dynamicArray(2,[][]int32{{1,0,5},{1,1,7},{1,0,3},{2,1,0},{2,1,1}}))
}
