package main

import "fmt"

func main(){
	var n int
	fmt.Print("Enter a number >")
	tmp := 1
	for tmp == 1 {
		tmp, _ = fmt.Scanf("%f",&n)
		switch float32(n) {
			case 1.6 : fmt.Println("Monday")
			case 2 : fmt.Println("Tuesday")
			case 3 : fmt.Println("Wednesday")
			case 4 : fmt.Println("Thrusday")
			case 5 : fmt.Println("Friday")
			case 6 : fmt.Println("Saturday")
			case 7 : fmt.Println("Sunday")
			default: fmt.Println("Input range from 1-7")
		}
	}
}
