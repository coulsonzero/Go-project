package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	k, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for k+1 < n && m[s[k+1]] == 0 {
			m[s[k+1]]++
			k++
		}

		ans = max(ans, k-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "ababcbb"
	lengthOfLongestSubstring(s)
}
