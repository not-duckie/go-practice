package main

import "fmt"
import "sync"

func printHello(i int){
	fmt.Println(i)
	wg.Done()
}
var wg = sync.WaitGroup{}
func main(){
	wg.Add(10)
	for i:=0;i<10;i++{
		go printHello(i)
	}
	wg.Wait()
}