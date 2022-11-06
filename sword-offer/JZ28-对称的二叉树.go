package main

/**
 * 判断一棵二叉树是不是对称的
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && check(root1.Left, root2.Right) && check(root1.Right, root2.Left)
}
