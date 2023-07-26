package listnode

/**
 * 要求：时间复杂度O(nlogN)
 * 输入：[{1,2},{1,4,5},{6}]
 * 输出：{1,1,2,4,5,6}
 */

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	h1, h2 := mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:])
	return merge(h1, h2)
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.Val < head2.Val {
		head1.Next = merge(head1.Next, head2)
		return head1
	} else {
		head2.Next = merge(head1, head2.Next)
		return head2
	}
}
