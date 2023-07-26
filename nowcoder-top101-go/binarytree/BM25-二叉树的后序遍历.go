package main

func postorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		order(node.Right)
		res = append(res, node.Val)
	}
	if root == nil {
		return nil
	}
	order(root)
	return res
}
