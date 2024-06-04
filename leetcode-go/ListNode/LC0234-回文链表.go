apackage main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 回文链表
// Input: head = [1,2,2,1]
// Output: true

// func isPalindrome(head *ListNode) bool {
// 	array := []int{head.Val}
//
// 	for ; head.Next != nil; head = head.Next {
// 		array = append(array, head.Next.Val)
// 	}
//
// 	for i := 0; i < len(array)/2; i++ {
// 		if array[i] != array[len(array)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }


func isPalindrome2(head *ListNode) bool {
	nums := []int{head.Val}
	for head.Next != nil {
		nums = append(nums, head.Next.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
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
