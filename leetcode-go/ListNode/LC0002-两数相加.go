package main

import (
	"fmt"
)

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

/**
 * 2. 两数相加
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 * 输入：l1 = [2,4,3], l2 = [5,6,4,2]
 * 输出：[7,0,8,2]
 * 解释：342 + 2465 = 2807.
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addList(l1, l2, 0)
}

func addList(l1 *ListNode, l2 *ListNode, v int) *ListNode {
	if l1 == nil && l2 == nil {
		if v == 0 {
			return nil
		}
		return &ListNode{Val: v, Next: nil}
	}
	if l1 != nil {
		v += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		v += l2.Val
		l2 = l2.Next
	}
	return &ListNode{v % 10, addList(l1, l2, v/10)}
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

	cur := addTwoNumbers(&l_node1, &r_node1)
	fmt.Println(cur)
	// fmt.Println(ToStringList(cur)) // 7 -> 0 -> 8 -> 2
}

// func ToStringList(cur *ListNode) (res string) {
// 	for cur != nil {
// 		if cur.Next != nil {
// 			res += strconv.Itoa(cur.Val) + "->"
// 		} else {
// 			res += strconv.Itoa(cur.Val)
// 		}
// 		cur = cur.Next
// 	}
// 	return res
// }
