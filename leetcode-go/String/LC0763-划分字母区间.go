package main

func partitionLabels(s string) []int {
	var res []int
	m := make(map[rune]int)

	for i, v := range s {
		m[v] = i
	}

	start := 0
	end := 0
	for i, v := range s {
		end = max(end, m[v])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
