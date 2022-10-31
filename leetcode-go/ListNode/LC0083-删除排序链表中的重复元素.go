package main

// 83. 删除排序链表中的重复元素
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
