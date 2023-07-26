package main

import "math"

/**
 * 幂：math.Pow(x, n) -> float64
 */

func printNumbers(n int) []int {
	var ans []int
	m := int(math.Pow10(n))
	for i := 1; i < m; i++ {
		ans = append(ans, i)
	}
	return ans
}
