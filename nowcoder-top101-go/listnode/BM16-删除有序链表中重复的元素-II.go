package listnode

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	if head.Val == next.Val {
		for next != nil && head.Val == next.Val {
			next = next.Next
		}
		return deleteDuplicates(next)
	} else {
		head.Next = deleteDuplicates(next)
		return head
	}
}
