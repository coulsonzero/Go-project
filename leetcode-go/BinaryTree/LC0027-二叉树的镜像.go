package main

/**
 * 剑指 Offer 27. 二叉树的镜像
 * 难度：简单
 * 描述：操作给定的二叉树，将其变换为源二叉树的镜像。
 * 输入：root = [4,2,7,1,3,6,9]
 * 输出：[4,7,2,9,6,3,1]
 */

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = mirrorTree(root.Left)
	root.Right = mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
