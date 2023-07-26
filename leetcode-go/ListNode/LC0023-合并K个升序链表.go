package main

import "sort"

/**
 * 23. 合并K个升序链表
 * 给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 * 输入：lists = []
 * 输出：[]
 * 输入：lists = [[]]
 * 输出：[]
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var ans []int
	for _, node := range lists {
		for node != nil {
			ans = append(ans, node.Val)
			node = node.Next
		}
	}
	sort.Ints(ans)

	// fmt.Println(ans)
	dummy := &ListNode{Next: nil}
	cur := dummy
	for _, v := range ans {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return dummy.Next
}
