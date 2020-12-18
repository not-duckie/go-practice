package main

import ("fmt")


func main(){
	a := 45
	var b *int = &a
	fmt.Printf("%p %d",b,*b)

}
