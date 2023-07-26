package 递归

/**
 * BM55 没有重复项数字的全排列
 * 难度：中等
 * 要求：空间复杂度O(n!), 时间复杂度O(n!)
 * 输入：[1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 */

func permute(nums []int) [][]int {
	// write code here
	res := make([][]int, 0)
	res = dfs(0, []int{}, nums, res)
	return res
}

func dfs(pos int, ans, nums []int, res [][]int) [][]int {
	if len(ans) == len(nums) {
		res = append(res, ans)
	}

	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]
		res = dfs(pos+1, append(ans, nums[pos]), nums, res)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
	return res
}
