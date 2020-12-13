package main

import ("fmt")

func main(){

	hashtable := make(map[string]int)

	hashtable = map[string]int {
		"Adnaman & Nicobar Islands": 744,
		"Anddra Pradesh": 50,
		"Arunachal Pradesh": 79,
		"Assam": 78,
		"Delhi": 11,
	}

	m := map[[3]int]string{}
	fmt.Print(hashtable,m)
}
