package main

import "fmt"

func divide(a,b int)(int, error){
	if b==0{
		return 0,fmt.Errorf("b cannot be zero")
	}

	return a/b, nil
}

func main(){
	res,err := divide(1,0)
	fmt.Errorf("fuck")
	fmt.Println(res,err)

}
