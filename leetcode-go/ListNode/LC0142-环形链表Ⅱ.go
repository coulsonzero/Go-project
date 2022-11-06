package main

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for ; head != nil; head = head.Next {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = head.Val
	}
	return head
}
