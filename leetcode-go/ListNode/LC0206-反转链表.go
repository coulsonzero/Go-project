package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func main() {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	head := ListNode{1, &node2}

	cur := reverseList(&head)
	fmt.Println(toString(cur)) // 3->2->1
}

func toString(cur *ListNode) (res string) {
	for cur != nil {
		if cur.Next != nil {
			res += strconv.Itoa(cur.Val) + "->"
		} else {
			res += strconv.Itoa(cur.Val)
		}
		cur = cur.Next
	}
	return res
}
