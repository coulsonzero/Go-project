package main

/*
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
*/

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	root := &TreeNode{Val: nums[m]}
	root.Left = helper(nums, l, m-1)
	root.Right = helper(nums, m+1, r)
	return root
}
