package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 回文链表
// Input: head = [1,2,2,1]
// Output: true

func isPalindrome(head *ListNode) bool {
	var array []int
	array = append(array, head.Val)
	for head.Next != nil {
		array = append(array, head.Next.Val)
		head = head.Next
	}
	// for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
	// 	if array[i] != array[j] {
	// 		return false
	// 	}
	// }
	for i := 0; i < len(array)/2; i++ {
		if array[i] != array[len(array)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	node4 := ListNode{1, nil}
	node3 := ListNode{2, &node4}
	node2 := ListNode{2, &node3}
	head := ListNode{1, &node2}
	println(isPalindrome(&head))
}
