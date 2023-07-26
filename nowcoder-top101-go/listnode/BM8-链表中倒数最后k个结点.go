package listnode

/**
 * BM8 链表中倒数最后k个结点
 * 输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
 * 如果该链表长度小于k，请返回一个长度为 0 的链表。
 * 输入：{1,2,3,4,5},2
 * 输出：{4,5}
 * 说明：返回倒数第2个节点4，系统会打印后面所有的节点来比较。
 */

func FindKthToTail(head *ListNode, k int) *ListNode {
	// write code here
	left, right := head, head
	for i := 0; i < k; i++ {
		if right == nil {
			return nil
		}
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}

	return left
}
