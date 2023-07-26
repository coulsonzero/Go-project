package main

/**
 * 6247. 从链表中移除节点
 * 难度：中等
 * 给你一个链表的头节点 head 。
 * 对于列表中的每个节点 node ，如果其右侧存在一个具有 严格更大 值的节点，则移除 node 。
 * 返回修改后链表的头节点 head 。
 * 输入：head = [5,2,13,3,8]
 * 输出：[13,8]
 * 解释：需要移除的节点是 5 ，2 和 3
 * 输入：head = [1,1,1,1]
 * 输出：[1,1,1,1]
 * 解释：每个节点的值都是 1 ，所以没有需要移除的节点。
 */

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := removeNodes(head.Next)
	if pre.Val <= head.Val {
		head.Next = pre
		return head
	} else {
		return pre
	}
}
