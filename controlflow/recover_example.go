package main

import ("fmt"
	"log")

func main(){
	fmt.Println("start")
	paniker()
	fmt.Println("stop")
}

func paniker(){
	fmt.Println("This is not good, i will panic")
	defer func(){
		if err := recover(); err != nil {
			log.Println("Error!!!:",err)
		}
	}()
	panic("WTF BRUH!!!")

}
