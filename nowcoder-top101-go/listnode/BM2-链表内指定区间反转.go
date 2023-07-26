package listnode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	dummyHead := &ListNode{Next: head}

	pre := dummyHead
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := m; i < n; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return dummyHead.Next
}
