package main

import ("fmt"
	"io"
	"net")

func main(){

	ln,err := net.Listen("tcp",":9000")
	if err!=nil{
		panic("Error Listening on port ")
	}
	fmt.Println("Started Listening")
	defer ln.Close()
	for {
		conn,err := ln.Accept()
		if err != nil{
			fmt.Println("Client Dropped due to error",err)
		}
		io.WriteString(conn,fmt.Sprint("bitch\n"))
	}
}

