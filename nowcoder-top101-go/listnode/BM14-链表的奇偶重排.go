package listnode

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	event, eventHead := head.Next, head.Next
	odd := head

	for event != nil && event.Next != nil {
		odd.Next = event.Next
		odd = odd.Next
		event.Next = odd.Next
		event = event.Next
	}
	odd.Next = eventHead
	return head
}
