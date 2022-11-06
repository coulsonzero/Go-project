package main

/**
 * 输入：head = [4,5,1,9], node = 5
 * 输出：[4,1,9]
 */

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
