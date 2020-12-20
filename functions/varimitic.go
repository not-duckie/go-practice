package main

import ("fmt")

func sum(num ...interface{}){
	fmt.Println(num)
}

func main(){
	sum(1,2,3,4,"wow",6.96)
}
