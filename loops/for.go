package main

import "fmt"


/*
*Go has only one loop "for" and it can be used for all types of loops
*/

func main(){
	for i:=0;i<5;i++{
		fmt.Println("The loop time #",i)
	}
	j :=0
	for j<5 {
		fmt.Println("The while loop time #",j)
		j++
	}
	k:=0
	for true {
		if k==10 { break }
		fmt.Println("the infinte loop ",k)
		k++
	}

}
