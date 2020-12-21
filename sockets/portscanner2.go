package main

import "fmt"
import "sync"
import "net"
import "strconv"


func scan(wg *sync.WaitGroup,port chan int,res chan int){
	for i:=range port{
		conn,err:= net.Dial("tcp","scanme.nmap.org:"+strconv.Itoa(i))
		if err==nil{
			res <- i
			conn.Close()
		}
		wg.Done()
	}
}

func printstuff(res chan int,wg *sync.WaitGroup){
	for i:=range res{
		fmt.Println(i)
	}
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	port := make(chan int, 50)
	rest := make(chan int, 1024)
	for i:=0;i<cap(port);i++{
		go scan(&wg,port,rest)
	}
	for i:=1;i<1024;i++{
		wg.Add(1)
		port <- i
	}
	wg.Add(1)
	go printstuff(rest,&wg)
	wg.Wait()
	defer close(port)
	defer close(rest)
}
