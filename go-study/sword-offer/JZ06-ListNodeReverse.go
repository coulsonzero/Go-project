package main

import "fmt"

/**
 * Q: JZ6-链表反转
 * input : {1,2,3}   1->2->3
 * output: {3,2,1}   1<-2<-3
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

var res []int

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return res
	}
	recursion(head)
	return res
}

func recursion(head *ListNode) {
	if head.Next != nil {
		recursion(head.Next)
	}
	res = append(res, head.Val)
}

func main() {
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	head := ListNode{Val: 1, Next: &node2}
	fmt.Println(printListFromTailToHead(&head)) // [3, 2, 1]
}
