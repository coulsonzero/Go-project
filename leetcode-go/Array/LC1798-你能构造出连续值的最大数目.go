package main

/*
给你一个长度为 n 的整数数组 coins ，它代表你拥有的 n 个硬币。第 i 个硬币的值为 coins[i] 。
如果你从这些硬币中选出一部分硬币，它们的和为 x ，那么称，你可以 构造 出 x 。
请返回从 0 开始（包括 0 ），你最多能 构造 出多少个连续整数。你可能有多个相同值的硬币。

输入：coins = [1,1,1,4]
输出：8
解释：你可以得到以下这些值：
- 0：什么都不取 []
- 1：取 [1]
- 2：取 [1,1]
- 3：取 [1,1,1]
- 4：取 [4]
- 5：取 [4,1]
- 6：取 [4,1,1]
- 7：取 [4,1,1,1]
从 0 开始，你可以构造出 8 个连续整数。
*/

import "sort"

func getMaximumConsecutive(coins []int) int {
	ans := 1
	sort.Ints(coins)
	for _, v := range coins {
		if v > ans {
			break
		}
		ans += v
	}
	return ans
}
