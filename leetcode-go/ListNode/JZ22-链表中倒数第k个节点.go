package main

/**
 * 给定一个链表: 1->2->3->4->5, 和 k = 2.
 * 返回链表 4->5.
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 统计链表长度
	cnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}

	// 倒数第K个：遍历n-k次即可
	pre := head
	for i := 0; i < cnt-k; i++ {
		pre = pre.Next
	}

	return pre
}

// 快慢指针
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Next
}
