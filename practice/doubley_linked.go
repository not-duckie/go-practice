package main

import "fmt"

type node struct{
	prev *node
	next *node
	data int
}

func pop(head *node,num int) int{
	var tmp1, tmp2 *node
	for head!=nil{
		if head.data == num{
			if head.next == nil {
				tmp1 = head.prev
				tmp1.next = nil
				return 0
			}
			tmp1 = head.prev
			tmp2 = head.next

			tmp1.next = tmp2
			tmp2.prev = tmp1
			return 0
		}
		head = head.next
	}
	fmt.Println("Element not in stack")
	return 1
}

func newnode(num int,head *node) *node{
	tmp := &node{
		data: num,}
	
	if(head==nil){
		tmp.next = nil
		head = tmp
		tmp.prev = nil
	} else {
		tmp.prev = nil
		tmp.next = head
		head.prev = tmp
		head = tmp
	}
	return head
}

func foredump(head *node){
	for head!=nil{
		fmt.Println(head.data)
		head = head.next
	}
}

func main(){
	var head *node
	head = newnode(5,head)
	head = newnode(10,head)
	head = newnode(15,head)
	head = newnode(20,head)
	foredump(head)
	pop(head,5)
	foredump(head)
}