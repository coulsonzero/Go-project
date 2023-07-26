package main

/**
 * 328. 奇偶链表
 * 难度：中等
 * 描述：给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
 * 		第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
 * 		请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
 * 		你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题
 * 输入: head = [1,2,3,4,5]
 * 输出: [1,3,5,2,4]
 * 输入: head = [2,1,3,5,6,4,7]
 * 输出: [2,3,6,7,1,5,4]
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	o := head      // 奇链表尾节点
	p := head.Next // 偶链表头结点
	e := p         // 偶链表尾节点

	for e != nil && e.Next != nil {
		o.Next = e.Next
		o = o.Next
		e.Next = o.Next
		e = e.Next
	}
	o.Next = p
	return head
}
