package main

/**
 * 给定一个头结点为 head 的非空单链表，返回链表的中间结点。
 * 如果有两个中间结点，则返回第二个中间结点
 * 输入：[1,2,3,4,5,6]
 * 输出：此列表中的结点 4 (序列化形式：[4,5,6])
 * 由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点
 */

// 快慢指针法: 时间复杂度：O(N) 空间复杂度：O(1)
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 计数
func middleNode2(head *ListNode) *ListNode {
	var length int
	for cur := head; cur.Next != nil; cur = cur.Next {
		length++
	}

	for i := 0; i < length/2; i++ {
		head = head.Next
	}

	if length%2 != 0 {
		return head.Next
	}
	return head
}

func middleNode3(head *ListNode) *ListNode {
	var length int
	cur := head
	for ; cur != nil; cur = cur.Next {
		length++
	}

	cur = head
	for i := 0; i < length/2; i++ {
		cur = cur.Next
	}

	return cur
}

// 数组法
func middleNode4(head *ListNode) *ListNode {
	var arr []*ListNode
	var t int
	for ; head != nil; head = head.Next {
		arr = append(arr, head)
		t++
	}
	return arr[t/2]

}
