package main
import "fmt"


func main(){
	fmt.Printf("The numbers chosen are 45 and 85")
	fmt.Printf(">")
	var ch rune
	tmp := 1
	for tmp == 1 {
		tmp, _ = fmt.Scanf("%c",&ch)
		if ch == '\n'{
		} else if ch == '+'{
			fmt.Println("The sum is ",45 + 85)
		} else if ch == '-' {
			fmt.Println("The subtraction result is ",45 - 85)
		} else if ch == '*' {
			fmt.Println("The product is ",45*85)
		} else if ch == '/' {
			fmt.Println("The division is ",float32 (45)/85)
		} else { fmt.Println("Enter +,-,* or /") }
	}
}
