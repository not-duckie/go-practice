package main

import ("fmt"
		"sync")

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func printHello(i int){
	fmt.Printf("fuckme %d\n",i)
	m.Unlock()
	wg.Done()
}

func main(){
	for i:=0;i<10;i++{
		wg.Add(1)
		m.Lock()
		go printHello(i)
	}
	wg.Wait()
}

