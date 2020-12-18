package main

import ("fmt"
	"strconv")

func countSheep(num int)(out string){
	for i:=1;i<=num;i++{
		out += strconv.Itoa(i) + " sheep..."
	}
	return out
}

func main(){
	fmt.Println(countSheep(3))

}
