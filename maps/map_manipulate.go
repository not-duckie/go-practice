package main

import "fmt"

func main(){
	hashmap := map[string]int{
		"Adaman & Nicobar Islands":74,
		"Andra Pradesh": 50,
		"Arunachal Pradesh":79,
		}
	//print specific index
	fmt.Println(hashmap["Andra Pradesh"])

	//Adding value to hashmap
	hashmap["Delhi"] = 11
	hashmap["sex"] = 11

	//printing the entire hashmap
	fmt.Println(hashmap)

	//Deleting an entry from map
	delete(hashmap,"sex")

	//printing the entire hashmap
	fmt.Println(hashmap)


	//Golang retunr the value as zero if the key doesnt exists
	// in the map. Thus 'ok' can be used to check if the key
	// exists or not or if the the actual value is zero.

	_, ok := hashmap["asdf"]
	fmt.Println(ok)


	// getting len of map
	fmt.Println(len(hashmap))

	//maps are passed by reference
	sp := hashmap
	delete(sp,"Delhi")
	fmt.Println(sp,"\n",hashmap)
	
	//printing all keys in map
	var a []string
	for key := range hashmap {
		fmt.Println("Key:",key)
		a = append(a,key)
	}
	fmt.Println(a)
}
