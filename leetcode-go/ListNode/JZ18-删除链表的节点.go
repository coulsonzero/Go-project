package main

/**
 * 剑指 Offer 18. 删除链表的节点(无重复元素)
 * 输入: head = [4,5,1,9], val = 5
 * 输出: [4,1,9]
 * 输入: head = [4,5,1,9], val = 1
 * 输出: [4,5,9]
 */

func deleteNode(head *ListNode, val int) *ListNode {
	// 删除头节点
	if head.Val == val {
		return head.Next
	}
	dummy := head
	for head != nil && head.Next != nil {
		// 删除尾节点
		if head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	return dummy
}

// 递归
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}
