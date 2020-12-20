package main

import "fmt"

type greeter struct{
	greeting string
	name string
}

func (g greeter) greet(){
	fmt.Println(g.greeting, g.name)
}

func (g greeter) ryhme1(){
	fmt.Println("ek do..ek do teen chaar")
	fmt.Println("peter ki gand maro yaar")
}

func test(g *greeter){
	fmt.Println(g.name)

}

func (g *greeter)test2(){
	g.name = "goel"

}

func main(){
	w := greeter{
		greeting: "Hello",
		name: "Go",
	}
	w.greet()
	peter := greeter {
			greeting: "peter",
			name: "bdsk",
			}
	peter.greet()
	peter.ryhme1()
	peter.greet()
	test(&peter)
	peter.test2()
	fmt.Println(peter.name)
}
