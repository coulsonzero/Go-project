package main

/**
 * 1171. 从链表中删去总和值为零的连续节点
 * 给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
 * 删除完毕后，请你返回最终结果链表的头节点。
 * 输入：head = [1,2,-3,3,1]
 * 输出：[3,1]
 * 提示：答案 [1,2,1] 也是正确的。
 * 输入：head = [1,2,3,-3,4]
 * 输出：[1,2,4]
 * 输入：head = [1,2,3,-3,-2]
 * 输出：[1]
 */

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sum := 0
	p := head
	for p != nil {
		sum += p.Val
		p = p.Next
		if sum == 0 {
			return removeZeroSumSublists(p)
		}
	}
	head.Next = removeZeroSumSublists(head.Next)
	return head
}
