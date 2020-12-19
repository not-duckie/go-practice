package main

import "fmt"
import "regexp"

func matchingStrings(strings []string, queries []string)(res []int32){
	for _,i := range queries {
		//regx := "^ *" + i + " *$"
		fmt.Println(i)
		regx := "^ ab $"
		if a,_ := regexp.MatchString(regx,strings[0]);a{
			fmt.Println(a)
		}
	}

	return []int32{0}
}

func main(){
	matchingStrings([]string{"ab"," ab","abc"},[]string{"ab","abc","bc"})

}
