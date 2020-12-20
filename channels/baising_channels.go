package main

import ("fmt"
		"sync")

func recv_only(ch <-chan string){
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	wg.Done()
}

func send_only(ch chan<- string){
	ch <- "wtf is happening"
	ch <- "go is so amazing bruh"
	wg.Done()
}
var wg = sync.WaitGroup{}
func main(){
	ch := make(chan string)
	wg.Add(2)
	go send_only(ch)
	go recv_only(ch)
	wg.Wait()
}