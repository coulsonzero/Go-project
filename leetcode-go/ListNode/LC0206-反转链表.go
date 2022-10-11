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
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
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
