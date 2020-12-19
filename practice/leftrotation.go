package main

import "fmt"

func rotate(d int32,arr []int32)([]int32){
	var tmp int32
	for tmp!=d{
		tmp++
		arr = append(arr,arr[0])
		arr = arr[1:]
	}
	return arr

}


func main(){
	fmt.Println(rotate(4,[]int32{1,2,3,4,5}))
}
