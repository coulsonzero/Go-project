package main

/**
 * 给定二叉树的根节点 root ，返回所有左叶子之和。
 * 输入: root = [3,9,20,null,null,15,7]
 * 输出: 24
 * 解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *   3
 *  / \
 * 9  20
 *    / \
 *   15  7
 */

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l = root.Left.Val
	}

	return l + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
