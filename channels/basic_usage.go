package main
import ("fmt"
		"sync")

var wg = sync.WaitGroup{}
func main(){
	ch := make(chan []int)
	wg.Add(3)
	go func(){
		num := <- ch
		for i:= range num{
			for j:= range num{
				if num[i]<num[j]{
					num[i],num[j] = num[j],num[i]
				}
			}
		}
	ch <- num
	wg.Done()
	}()
	go func(){
		num := <- ch
		fmt.Println(num)
		wg.Done()
	}()
	go func(){
		ch <- []int{10,9,8,7,6,5,4,3,2,1}
		wg.Done()
	}()
	wg.Wait()
}