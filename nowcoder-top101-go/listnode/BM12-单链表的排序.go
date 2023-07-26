package listnode

import "sort"

func sortInList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return nil
	}
	temp := make([]int, 0)
	t1 := head
	for t1 != nil {
		temp = append(temp, t1.Val)
		t1 = t1.Next
	}
	sort.Ints(temp)
	t2 := head
	i := 0
	for t2 != nil {
		t2.Val = temp[i]
		t2 = t2.Next
		i++
	}
	return head
}
