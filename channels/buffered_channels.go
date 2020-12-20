package main

import "fmt"
import "sync"

var wg = sync.WaitGroup{}

func channel1(ch chan string){
	ch <- "Hey bruh"
	ch <- "whats up ?"
	close(ch)
	wg.Done()
}

func channel2(ch chan string){
	for i:=range ch{
		fmt.Println(string(i[0]))
	}
	wg.Done()
}

func main(){
	ch := make(chan string,50)
	wg.Add(2)
	go channel1(ch)
	go channel2(ch)
	wg.Wait()
}