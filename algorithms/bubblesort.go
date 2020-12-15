package main

import ("fmt")

func swap(a, b *int){
	*a = (*a) + (*b)
	*b = (*a) - (*b)
	*a = (*a) - (*b)
}

func main(){
	ar := [10]int{10,9,8,7,6,5,4,3,2,1}
	for i:=0;i<len(ar);i++{
		for j :=0;j<len(ar);j++{
			if (ar[i]<ar[j]){
				ar[j],ar[i] = ar[i],ar[j]
			}
		}
	}
fmt.Println(ar)
}
