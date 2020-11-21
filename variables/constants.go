package main

import "fmt"
import "math"

func main(){
	const s string = "constant string"

	const n = 5000000

	const d = 3e20 / n
	fmt.Println("The value of var d is ",int64(d))

	fmt.Printf("%3f",float32(math.Sin(45)))
}
