package main

import "fmt"

func main(){
	var a = "peter"
	var b = 69
	var c,d = 6.9,4
	fmt.Printf("The var a has type %T and value %s\n",a,a)
	fmt.Printf("The var b has type %T and value %d \n",b,b)
	fmt.Printf("The var c has type %T and value %f \n",c,c)
	fmt.Printf("The var d has type %T and value %d \n",d,d)

	f:="short hand"
	fmt.Println("This is shorthand for decalring variables ",f)
}
