package main

import ("fmt")

func takeinput(in [][3]int){
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			fmt.Scanf("%d",&in[i][j])
		}
	}
}

func output(out [][3]int){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++ {
			fmt.Printf("%d\t",out[i][j])
		}
		fmt.Println()
	}
}

func multiplication(a [][3]int, b [][3]int, c[][3]int){
	sum :=0
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			for k:=0;k<3;k++ {
				sum = sum + a[i][k]*b[k][j]
			}
		c[i][j] = sum
		sum = 0
		}
	}
}

func main(){
	var a[3][3]int
	var b[3][3]int
	var c[3][3]int

	fmt.Println("Enter the value for matrix one")
	takeinput(a[:])
	output(a[:])

	fmt.Println("Enter the value for matrix two")
	takeinput(b[:])
	output(b[:])

	fmt.Println()

	multiplication(a[:],b[:],c[:])
	output(c[:])
}
