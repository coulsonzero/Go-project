package main

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	t, sum := skill[0]+skill[n-1], 0
	for i := 0; i < n; i++ {
		if skill[i]+skill[n-1-i] != t {
			return -1
		}
		sum += skill[i] * skill[n-1-i]
	}
	return int64(sum) / 2
}
