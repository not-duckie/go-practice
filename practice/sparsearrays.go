package main

import "fmt"
import "strings"



func matchingStrings(str []string, queries []string)(res []int32){
	for _,i := range queries{
		var count int32
		for _,j := range str{
			j = strings.ReplaceAll(j," ","")
			if i == j{
				count++
			}
		}
		res = append(res,count)
	}
	return res
}

func main(){
	fmt.Println(matchingStrings([]string{"ab"," ab","abc"},[]string{"ab","abc","bc"}))

}
