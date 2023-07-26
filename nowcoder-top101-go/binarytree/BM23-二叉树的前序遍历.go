package main

func preorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		order(root.Left)
		order(root.Right)
	}

	if root == nil {
		return nil
	}
	order(root)
	return res
}
