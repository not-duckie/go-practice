package main

import "fmt"

func main(){
	fmt.Printf("Enter a number >")

	for true {
		var n int
		fmt.Scanf("%d",&n)

			if n%2==0{
				fmt.Printf("%d is even\n",n)
			} else { fmt.Println("not even") }
	}
}
