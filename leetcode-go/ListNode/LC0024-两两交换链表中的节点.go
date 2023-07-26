package main

/*
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 * 输入：head = []
 * 输出：[]
 * 输入：head = [1]
 * 输出：[1]
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head.Next
	head.Next = swapPairs(cur.Next)
	cur.Next = head
	return cur
}

// 切片存储
func swapPairs2(head *ListNode) *ListNode {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	// fmt.Printf("%v\n", nums)

	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	// fmt.Printf("%v\n", nums)
	dummy := &ListNode{Next: nil}
	cur := dummy
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return dummy.Next
}
