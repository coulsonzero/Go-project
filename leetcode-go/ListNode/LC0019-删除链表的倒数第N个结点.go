package main

/**
 * Question: 删除链表的倒数第 N 个结点
 * Descript: 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 暴力算法：先获取链表长度，然后遍历到倒数第n个节点后删除改节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	getLength := func(head *ListNode) int {
		var length int
		for head != nil {
			head = head.Next
			length++
		}
		return length
	}

	length, dummy := getLength(head), &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 快慢指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 当快指针移出链表时(fast.Next == nil), 慢指针删除该节点
	slow.Next = slow.Next.Next
	return head
}
