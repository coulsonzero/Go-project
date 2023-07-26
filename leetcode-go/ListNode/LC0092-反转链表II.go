package main

/**
 * question: 92-反转链表II
 * descript: 反转链表中间部分节点 (给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置             right 的链表节点，返回 反转后的链表)
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 * 描述：从第二个节点开始反转 2->3->4 反转为 4->3->2, 连接头部节点和尾部节点, 1->4->3->2->5
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, pre.Next, cur.Next
	}
	return dummy.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 返回反转的链表节点和末尾节点
	cur, r := reverseList(pre.Next, left, right)

	pre.Next.Next = r
	pre.Next = cur

	return dummy.Next
}

func reverseList(head *ListNode, left, right int) (*ListNode, *ListNode) {
	var pre *ListNode
	for i := left; i <= right; i++ {
		head, head.Next, pre = head.Next, pre, head
	}
	return pre, head
}
