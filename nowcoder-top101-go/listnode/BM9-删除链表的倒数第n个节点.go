package listnode

/**
 * BM9 删除链表的倒数第n个节点
 * 给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
 * 输入：1 → 2 → 3 → 4 → 5    2
 * 输出：1 → 2 → 3 → 5.
 * 说明：删除倒数第2个节点
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
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
	slow.Next = slow.Next.Next
	return head
}
