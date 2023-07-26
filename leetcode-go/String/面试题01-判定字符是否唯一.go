package main

func isUnique(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
		if m[v] > 1 {
			return false
		}
	}

	return true
}
