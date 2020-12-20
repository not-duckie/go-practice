package main

import "fmt"

type ConsoleWriter struct {}

type fuck struct {}

type writer interface{
	Write([]byte) (int ,error)
}

func (cw ConsoleWriter) Write(data []byte)(int,error){
	n,err := fmt.Println(string(data))
	return n,err
}

func main(){
	var w ConsoleWriter = ConsoleWriter{}
	w.Write([]byte("hello bitch"))
}