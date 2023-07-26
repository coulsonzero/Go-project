package 递归

func permuteUnique(nums []int) [][]int {
	// write code here
	res := make([][]int, 0)
	res = dfs(0, []int{}, nums, res)
	return res
}

func dfs(pos int, ans, nums []int, res [][]int) [][]int {
	if len(ans) == len(nums) {
		res = append(res, ans)
	}

	m := make(map[int]bool)
	for i := pos; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = true

		nums[pos], nums[i] = nums[i], nums[pos]
		res = dfs(pos+1, append(ans, nums[pos]), nums, res)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
	return res
}
