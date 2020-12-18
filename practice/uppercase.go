package main


import "regexp"
import "fmt"

type MyString string

func (s MyString) IsUpperCase() bool {
	a,_:=regexp.MatchString("^[^a-z]$",string(s))
	return a
}

func main(){
	fmt.Println(MyString("{okay}").IsUpperCase())

}