package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := make(map[int]int)
	for head != nil {
		m[head.Val]++
		head = head.Next
	}

	// 删除重复节点
	nums := make([]int, 0)
	for k, _ := range m {
		if m[k] == 1 {
			nums = append(nums, k)
		}
	}

	// fmt.Printf("%v\n", nums)
	sort.Ints(nums)

	// 生成新的链表
	dummy := &ListNode{Next: nil}
	cur := dummy
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return dummy.Next

}
