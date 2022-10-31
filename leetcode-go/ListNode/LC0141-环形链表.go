package main

// map 哈希表
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func main() {
	var node1, node2, node3, node4 ListNode
	node1 = ListNode{3, &node2}
	node2 = ListNode{2, &node3}
	node3 = ListNode{0, &node4}
	node4 = ListNode{4, &node2}
	println(hasCycle(&node1))

}
