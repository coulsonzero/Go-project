package listnode

func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1，找链表中点，双数就是后面的数
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 2，翻转链表
	var pre *ListNode
	cur := slow
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	// 3，比较链表值大小
	mid := pre
	for mid != nil {
		if mid.Val != head.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}
	return true
}

func isPail2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	r := len(res) - 1
	for l := 0; l < r; l++ {
		if res[l] != res[r] {
			return false
		}
		r--
	}
	return true
}
