package main

import ("fmt"
	"math/rand"
	"time"
	"os")

func main(){
	src := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(src)
	num := r1.Intn(100)
	if (num<=50){
		fmt.Println("Hint#1 The number is less than or equal to 50")
	} else { fmt.Println("Hint#1 The number is greater than 50") }
	n := 0
	tmp := 0
	for n != num {
		tmp++
		fmt.Print(">")
		fmt.Scanf("%d",&n)
		if n>num {
			fmt.Println("Your guess is too big")
		} else if n<num { fmt.Println("Your guess it too small")
		} else { fmt.Println("You guessed it right !!")
			 fmt.Println("You guess it in ",tmp," tries")
			 os.Exit(0)
			}
	}
}
