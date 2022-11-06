package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m1 := minDepth(root.Left)
	m2 := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	}
	return min(m1, m2) + 1
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	switch {
	case root.Left == nil:
		return minDepth(root.Right) + 1
	case root.Right == nil:
		return minDepth(root.Left) + 1
	default:
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
