package main

import ("fmt")

func Solution(word string)(out string){
	for _,i := range word{
		out = string(i) + out
	}
	return out
}

func main(){
	fmt.Println(Solution("world😀"))
}
