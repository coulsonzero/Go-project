package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	// step1. 将链表节点添加到切片中
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	// step2. 每 k 个一组进行切片翻转, 少于k个的末尾元素不反转
	reverseNums := func(arr []int) {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
	}
	for i := 0; i < len(nums); i += k {
		if i+k <= len(nums) {
			reverseNums(nums[i : i+k])
		}
	}

	// step3. 将切片生成新的链表
	dummy := &ListNode{Next: nil}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return dummy.Next
}
