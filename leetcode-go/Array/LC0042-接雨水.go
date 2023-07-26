package main

func trap(height []int) int {
	l := 0
	r := len(height) - 1
	maxL := 0
	maxR := 0

	ans := 0
	for l < r {
		maxL = max(maxL, height[l])
		maxR = max(maxR, height[r])
		if maxL < maxR {
			ans += maxL - height[l]
			l++
		} else {
			ans += maxR - height[r]
			r--
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
