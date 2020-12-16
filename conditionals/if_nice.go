package main

import "fmt"

func main() {
	hashmap := map[string]int{"Anadman Nicobar islands": 744,
		"Andra Pradesh": 50}
	if pop,ok:=hashmap["Andra Pradesh"]; ok{
		fmt.Println("The Andra Pradesh has pincode starting with",pop)
	}
	fmt.Println(hashmap)
}
