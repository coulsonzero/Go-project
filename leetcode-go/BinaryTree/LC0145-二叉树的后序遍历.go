package main

// 二叉树的后序遍历
// Input: root = [1,null,2,3]
// Output: [3,2,1]

func postorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
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
	return
}
