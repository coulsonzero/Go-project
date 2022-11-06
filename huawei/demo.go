package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l_node1, l_node2, l_node3 ListNode
	l_node1 = ListNode{2, &l_node2}
	l_node2 = ListNode{4, &l_node3}
	l_node3 = ListNode{3, nil}

	var r_node1, r_node2, r_node3, r_node4 ListNode
	r_node1 = ListNode{5, &r_node2}
	r_node2 = ListNode{6, &r_node3}
	r_node3 = ListNode{4, &r_node4}
	r_node4 = ListNode{2, nil}

	addTwoNumbers(&l_node1, &r_node1)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	var head *ListNode
	cur := head
	var carry, v int
	for l1 != nil || l2 != nil {
		v = carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		carry, v = v%10, v/10
		cur = &ListNode{Val: v}
		cur.Next = cur
	}

	fmt.Println(head)
	fmt.Println(head.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next.Next)

	// var arr []int
	// if cur != nil {
	// 	arr = append(arr, cur.Val)
	// 	cur = cur.Next
	// }
	// // fmt.Println(cur)
	// fmt.Println(arr)
}

func toStringList(cur *ListNode) (res string) {
	for cur != nil {
		if cur.Next != nil {
			res += strconv.Itoa(cur.Val) + " -> "
		} else {
			res += strconv.Itoa(cur.Val)
		}
		cur = cur.Next
	}
	return res
}
