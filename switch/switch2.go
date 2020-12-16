package main

import "fmt"

func main(){
	var a int
	//only accepts integer similar to C language
	fmt.Print("Enter a number >")
	fmt.Scanf("%d",&a)
	//can use the keyword fallthrough to fall
	switch a {
	case 69,420: fmt.Println("nice")
	fallthrough
	case 56: fmt.Println("easter egg bitch") 
	default: fmt.Println("not nice")
	}

}