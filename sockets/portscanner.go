package main

import ("fmt"
	"net"
	"strconv"
	"sync")

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

/*
func(i int){
	conn,err:=net.Dial("tcp","scanme.nmap.org:"+strconv.Itoa(i))
	if err == nil{
		fmt.Println("The port is open",i,"; connection success")
		conn.Close()
	} else { fmt.Println("The port ",i," is closed") }
	wg.Done()
}
*/
func worker(ports chan int){
	for i:=range ports{
		conn,err:=net.Dial("tcp","scanme.nmap.org:"+strconv.Itoa(i))
		if err == nil{
			fmt.Println("The port is open",i,"; connection success")
			conn.Close()
		} else { fmt.Println("The port ",i," is closed") }
		wg.Done()
	}
}

func main(){
	ports := make(chan int,100)
	for i:=0;i<cap(ports);i++{
		go worker(ports)
	}
	for i:=1;i<=1024;i++{
		wg.Add(1)
		ports <-i
	}
	wg.Wait()
	defer close(ports)
}
