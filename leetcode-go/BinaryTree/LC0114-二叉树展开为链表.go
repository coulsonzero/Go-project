package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	// postorder
	left, right := root.Left, root.Right
	root.Left = nil
	root.Right = left
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
