package array

func maxLength(arr []int) int {
	// write code here
	res := 0
	start := 0
	m := make(map[int]int, 0)
	for i, v := range arr {
		index, ok := m[v]
		m[v] = i
		if ok && start <= index {
			start = index + 1
		}

		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}
