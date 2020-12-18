package main

import ("fmt")

func main(){
	fmt.Println("This normals thing")
	fmt.Println("This is not normal, pls panik")
	defer fmt.Println("end")
	panic("PANIC!!")

}
