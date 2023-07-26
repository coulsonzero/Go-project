package listnode

/**
 * 合并两个排序的链表
 * 要求：空间复杂度O(1), 时间复杂度O(n)
 * 输入：{1,3,5},{2,4,6}
 * 输出：{1,2,3,4,5,6}
 */

func Merge(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	if head1.Val < head2.Val {
		head1.Next = Merge(head1.Next, head2)
		return head1
	} else {
		head2.Next = Merge(head2.Next, head1)
		return head2
	}
}
