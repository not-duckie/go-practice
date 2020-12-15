package main


import (
"fmt"
"net/http"
)


func main(){
	resp,_ := http.Get("http://example.com/")
	for i := range resp{
		fmt.Println(i)
	}
}
