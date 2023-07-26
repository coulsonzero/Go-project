package listnode

func FindFirstCommonNode(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	h1, h2 := head1, head2
	if h1 == nil || h2 == nil {
		return nil
	}

	for h1 != h2 {
		if h1 == nil {
			h1 = head2
		} else {
			h1 = h1.Next
		}

		if h2 == nil {
			h2 = head1
		} else {
			h2 = h2.Next
		}
	}
	return h2
}
