package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// write code here
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}
