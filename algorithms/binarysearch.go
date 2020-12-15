package main

import ("fmt"
	"os")

func main(){
	ar := [...]int{10,9,8,7,6,5,4,3,2,1}
	var x int
	fmt.Printf("Enter the element to find\n")
	fmt.Scanf("%d",&x)
	low := 0
	high := 9
	for low < high {
		mid := (low+high)/2
		if(ar[mid]==x){
			fmt.Println("found the element ", x," at ", mid+1)
			os.Exit(0)
		}
		if (ar[mid] > x){
			low = mid
		}
		if (ar[mid]<x){
			high = mid
		}
	}
	fmt.Println("Sorry Could not found the element")
}
