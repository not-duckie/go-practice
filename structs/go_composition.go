package main

import "fmt"

type Animal struct {
	name string
	origin string
}

type Brid struct {
	Animal
	speed float32
	canfly bool

}

func main(){
	b := Brid{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speed = 48
	b.canfly = false
	fmt.Println(b)
	c := Brid{
		Animal: Animal{name: "peter",origin: "wonderland"},
		speed: 69,
		canfly: true,
	}
	fmt.Println(c)

}
