package main

func levelOrder(root *TreeNode) [][]int {
	// write code here
	var res [][]int
	var order func(node *TreeNode, level int)
	order = func(node *TreeNode, level int) {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			order(node.Left, level+1)
		}
		if node.Right != nil {
			order(node.Right, level+1)
		}
	}

	if root == nil {
		return nil
	}
	order(root, 0)
	return res
}
