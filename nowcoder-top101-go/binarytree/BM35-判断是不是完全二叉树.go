package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) {
	// write code here
	var order func(node *TreeNode, cnt int)
	order = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}

		fmt.Printf("node: %v, cnt: %v \n", node, cnt)
		order(node.Left, cnt+1)
		order(node.Right, cnt+1)

	}
	order(root, 1)
}

func main() {
	var head, node2, node3, node4, node5, node6, node7 *TreeNode
	head = &TreeNode{1, node2, node3}
	node2 = &TreeNode{2, node4, node5}
	node3 = &TreeNode{3, node6, node7}
	node4 = &TreeNode{4, nil, nil}
	node5 = &TreeNode{5, nil, nil}
	node6 = &TreeNode{6, nil, nil}
	node7 = &TreeNode{7, nil, nil}
	isCompleteTree(head)
}
