package main

import "fmt"

/**
 * Question: JZ32 从上往下打印二叉树
 * input: {8,6,10,#,#,2,1}
 * input: [8,6,10,2,1]
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	res, queue := []int{}, []*TreeNode{root}
	if root == nil {
		return res
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func main() {
	node5 := TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := TreeNode{Val: 2, Left: nil, Right: nil}
	node3 := TreeNode{Val: 10, Left: &node4, Right: &node5}
	node2 := TreeNode{Val: 6, Left: nil, Right: nil}
	root := TreeNode{Val: 8, Left: &node2, Right: &node3}
	fmt.Println(PrintFromTopToBottom(&root))
}
