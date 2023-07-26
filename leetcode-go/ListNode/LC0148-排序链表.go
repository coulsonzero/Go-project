package main

import "sort"

func sortList(head *ListNode) *ListNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	sort.Ints(nums)

	// 生成新链表
	dummy := &ListNode{Next: nil}
	cur := dummy
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return dummy.Next
}
