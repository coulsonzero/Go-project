package main

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] == 3 {
			delete(m, v)
		}
	}

	ans := 0
	for k, _ := range m {
		ans = k
	}
	return ans
}
