package main

/**
 * 剑指 Offer 34. 二叉树中和为某一值的路径
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 */

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var temp []int

	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		temp = append(temp, root.Val)
		if root.Left == nil && root.Right == nil && root.Val == target {
			res = append(res, append([]int{}, temp...))
		}
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
		temp = temp[:len(temp)-1]
	}
	dfs(root, target)
	return res
}
