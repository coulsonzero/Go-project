package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	sum := 0
	for j := 0; j < len(grid[0]); j++ {
		tmp := 0
		for i := 0; i < len(grid); i++ {
			tmp = max(grid[i][j], tmp)
		}
		sum += tmp
	}

	return sum
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
