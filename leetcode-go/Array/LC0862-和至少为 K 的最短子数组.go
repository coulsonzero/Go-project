package main

/**
 * 862. 和至少为 K 的最短子数组
 * 难度：困难
 * 给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1
 * 输入：nums = [1,2], k = 4
 * 输出：-1
 * 输入：nums = [2,-1,2], k = 3
 * 输出：3
 */

// 优先队列
func shortestSubarray(nums []int, k int) int {
	n, sum, q := len(nums), 0, [][]int{}
	res := n + 1
	q = append(q, []int{0, -1})
	for i := 0; i < n; i++ {
		sum += nums[i]
		for len(q) > 0 && sum <= q[len(q)-1][0] {
			q = q[:len(q)-1]
		}
		for len(q) > 0 && sum-q[0][0] >= k {
			res, q = min(res, i-q[0][1]), q[1:]
		}
		q = append(q, []int{sum, i})
	}
	if res == n+1 {
		return -1
	}
	return res
}

/*
执行结果：通过
执行用时：160 ms, 在所有 Go 提交中击败了 20.79% 的用户
内存消耗：9 MB, 在所有 Go 提交中击败了 87.53% 的用户
*/

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

/*
// 超出时间限制
func shortestSubarray(nums []int, k int) int {
	var temp []int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if sum(nums[i:j+1]) >= k {
				temp = append(temp, len(nums[i:j+1]))
			}
		}
	}
	if len(temp) == 0 {
		return -1
	}
	return minNum(temp)
}

func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}

func minNum(nums []int) int {
	res := nums[0]
	for _, v := range nums {
		if v <= res {
			res = v
		}
	}
	return res
}
*/
