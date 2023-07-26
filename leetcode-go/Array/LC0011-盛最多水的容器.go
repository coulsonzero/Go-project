package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		minH := min(height[l], height[r])
		area := minH * (r - l)
		ans = max(ans, area)

		for height[l] <= minH && l < r {
			l++
		}
		for height[r] <= minH && l < r {
			r--
		}
	}

	return ans
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
//
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 1}))
}
