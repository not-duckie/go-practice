package main

import ("fmt")

func insertionsort(ar [10]int) [10]int {
	for i:=1;i<len(ar);i++{
		key := ar[i]
		j := i-1
		for j>=0 && ar[j]>key{
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = key
	}

	return ar
}

func main(){
	ar := [...]int{10,9,8,7,6,5,4,3,2,1}
	ar = insertionsort(ar)
	fmt.Println(ar)

}
