package main

import (
	"strconv"
)

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// insertNode 插入节点
func insertNode(node *ListNode, k *ListNode) {
	// temp := node.Next
	// node.Next = k
	// k.Next = temp
	node.Next, k.Next = k, node.Next
}

// delNode 移除链表指定节点
func delNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// deleteDuplicates 移除重复节点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// toStringList 打印链表
func toStringList(head *ListNode) (res string) {
	for ; head != nil; head = head.Next {
		if head.Next != nil {
			res += strconv.Itoa(head.Val) + "->"
		} else {
			res += strconv.Itoa(head.Val)
		}
	}
	return res
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

// isPalindrome 判断是否回文
func isPalindrome(head *ListNode) bool {
	array := []int{head.Val}

	for ; head.Next != nil; head = head.Next {
		array = append(array, head.Next.Val)
	}

	for i := 0; i < len(array)/2; i++ {
		if array[i] != array[len(array)-1-i] {
			return false
		}
	}
	return true
}

// hasCycle 判断环形链表
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for ; head != nil; head = head.Next {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = head.Val
	}
	return false
}

func main() {
	var head, node2, node3, node4 ListNode
	head = ListNode{4, &node2}
	node2 = ListNode{5, &node3}
	node3 = ListNode{1, &node4}
	node4 = ListNode{9, nil}

	println(toStringList(&head))

	// delNode(&node2)
	var k ListNode
	k.Val = 7
	insertNode(&node2, &k)
	println(toStringList(&head))
}
