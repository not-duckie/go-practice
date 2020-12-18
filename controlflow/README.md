# Defer Functions

Consider the following code,

```
func main(){
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

}

```

Now the control flow happens this way, first the start statement is printed and
the go see the next is defer statement so it leaves it and then prints the end
statment and then the main function returns and now go checks if there were any
defer statments and now it executes them.
