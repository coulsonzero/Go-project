package main

func findContinuousSequence(target int) [][]int {
	// (r - l + 1) * (l + r) / 2  = target

	var ans [][]int
	l := 1
	r := 2
	for l < r {
		sum := (r - l + 1) * (l + r) / 2
		if sum == target {
			tmp := make([]int, r-l+1)
			for i := 0; i < r-l+1; i++ {
				tmp[i] = l + i
			}
			ans = append(ans, tmp)
			l++
		} else if sum < target {
			r++
		} else if sum > target {
			l++
		}
	}

	return ans
}
