package main

import "fmt"

type node struct{
	ptr *node
	data int
}

func newnode(num int, head *node) *node {
	var tmp *node
	if(head == nil){
		tmp = &node{
			ptr: nil,
			data: num,
		}
	} else {
		tmp = &node{
			ptr: head,
			data: num,
		}
	}
	head = tmp
	return head
}

func dumplist(head *node){
	for head!=nil {
		fmt.Println(head.data)
		head = head.ptr
	}
}

func main(){
	var head *node
	head = newnode(45,head)
	fmt.Println(head)
	head = newnode(85,head)
	fmt.Println(head)
	head = newnode(20,head)
	fmt.Println(head)
	dumplist(head)

}