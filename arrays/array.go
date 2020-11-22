package main

import (
"fmt"
)

func main(){
	var a[10]int
	fmt.Println("emp:",a)
	for i:=0;i<10;i++{
		a[i] = i
	}
	fmt.Println("filled",a)

}
