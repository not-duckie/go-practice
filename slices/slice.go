package main

import ("fmt")


func fast_pop(a []string,i int){
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Printf("%T\n",a)

}

func main(){
	a := []string{"A","B","C","D"}
	i := 1
	fast_pop(a,i)
	for i := range a  {
		fmt.Println(a[i])
	}
}
