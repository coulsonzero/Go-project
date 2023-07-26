package main

func longestSquareStreak(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	ans := 0
	for k := range m {
		cnt := 1
		for k *= k; m[k]; k *= k {
			cnt++
		}
		ans = max(ans, cnt)
	}

	if ans == 1 {
		return -1
	}

	return ans
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
