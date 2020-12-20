package main

import ("fmt")

type ListNode struct{
	Val int
	Next *ListNode
}

func newnode(head *ListNode, val int)(*ListNode){
    var tmp *ListNode
    tmp = &ListNode{}
    tmp.Val = val
    if(tmp==nil){
        tmp.Next = nil
        head = tmp
    } else {
        tmp.Next = head
        head = tmp
    }
    return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var tmp *ListNode
    var car int
    for l1!=nil && l2!=nil{
        sum := l1.Val + l2.Val + car
        if sum >10{
            sum = sum%10
            fmt.Printf("LOG:",sum)
            car = 1
        } else {
            car = 0
            }
    tmp = newnode(tmp,sum)
    l1 = l1.Next
    l2 = l2.Next
    }
    return (tmp)
}

func dumplist(head *ListNode){
	for head!=nil{
		fmt.Println(head.Val)
		head = head.Next
	}
}


func main(){
	var list1 *ListNode
	var list2 *ListNode

	list1 = newnode(list1,2)
	list1 = newnode(list1,4)
	list1 = newnode(list1,3)

	list2 = newnode(list2,5)
	list2 = newnode(list2,6)
	list2 = newnode(list2,4)

	tmp := addTwoNumbers(list1,list2)
	dumplist(tmp)
}
