package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 != nil { h1 = h1.Next } else { h1 = headB }
		if h2 != nil { h2 = h2.Next } else { h2 = headA }
	}

	return h1
}
