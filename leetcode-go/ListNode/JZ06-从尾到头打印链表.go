package main

import "fmt"

/**
 * 输入：head = [1,3,2]
 * 输出：[2,3,1]
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reversePrint(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	reverse := func(nums []int) {
		r := len(nums) - 1
		for l := 0; l < r; l++ {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}

	reverse(ans)
	return ans
}

func main() {
	var head, node2, node3 *ListNode
	head = &ListNode{1, node2}
	node2 = &ListNode{2, node3}
	node3 = &ListNode{3, nil}
	fmt.Printf("res: %v \n", reversePrint(head))
}
