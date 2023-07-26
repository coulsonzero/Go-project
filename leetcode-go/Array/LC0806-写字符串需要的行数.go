package main

func numberOfLines(widths []int, s string) []int {
	ans := make([]int, 2)
	ans[0] = 1
	for _, v := range s {
		ans[1] += widths[v-'a']
		if ans[1] > 100 {
			ans[1] = widths[v-'a']
			ans[0]++
		}
	}

	return ans
}
