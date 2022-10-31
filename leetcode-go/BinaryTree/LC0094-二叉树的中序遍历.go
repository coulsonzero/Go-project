package main

// 二叉树的中序遍历
// Input: root = [1,null,2,3]
// Output: [1,3,2]

func inorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		res = append(res, node.Val)
		order(node.Right)
	}
	order(root)
	return
}
