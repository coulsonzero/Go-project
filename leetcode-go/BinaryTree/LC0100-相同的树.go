package main

/**
 * 100. 相同的树
 * 难度：简单
 * 输入：p = [1,2,3], q = [1,2,3]
 * 输出：true
 * 输入：p = [1,2], q = [1,null,2]
 * 输出：false
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
