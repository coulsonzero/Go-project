package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// [3,9,20,null,null,15,7]
	node5 := TreeNode{7, nil, nil}
	node4 := TreeNode{15, nil, nil}
	node3 := TreeNode{20, &node4, &node5}
	node2 := TreeNode{9, nil, nil}
	node1 := TreeNode{3, &node2, &node3}

	println(maxDepth(&node1)) // 3
}
