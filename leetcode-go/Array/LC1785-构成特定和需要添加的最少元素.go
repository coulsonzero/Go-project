package main

/**
 * 1785. 构成特定和需要添加的最少元素
 * 给你一个整数数组 nums ，和两个整数 limit 与 goal 。数组 nums 有一条重要属性：abs(nums[i]) <= limit 。
 * 返回使数组元素总和等于 goal 所需要向数组中添加的 最少元素数量 ，添加元素 不应改变 数组中 abs(nums[i]) <= limit 这一属性。
 * 注意，如果 x >= 0 ，那么 abs(x) 等于 x ；否则，等于 -x 。
 * 输入：nums = [1,-1,1], limit = 3, goal = -4
 * 输出：2
 * 解释：可以将 -2 和 -3 添加到数组中，数组的元素总和变为 1 - 1 + 1 - 2 - 3 = -4 。
 * 输入：nums = [1,-10,9,1], limit = 100, goal = 0
 * 输出：1
 */

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	a := abs(goal-sum) / limit
	b := abs(goal-sum) % limit

	// fmt.Println(a, b)
	if b == 0 {
		return a
	}
	return a + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
