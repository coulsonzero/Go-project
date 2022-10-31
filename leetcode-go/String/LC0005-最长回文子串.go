package main

/**
 * Question: 5. 最长回文子串
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案
 * 输入：s = "cbbd"
 * 输出："bb"
 */

func longestPalindrome(s string) string {
	var res string
	for i, l, r := 0, 0, 0; i < len(s); i++ {
		l, r = i, i
		for r+1 < len(s) && s[r+1] == s[i] {
			r++
		}
		for l-1 >= 0 && r+1 < len(s) && s[r+1] == s[l-1] {
			r++
			l--
		}
		if r > l && len(res) < r-l+1 {
			res = s[l : r+1]
		}
	}
	if len(res) == 0 {
		return string(s[0])
	}
	return res
}

/*
执行结果：通过
执行用时：4 ms, 在所有 Go 提交中击败了 93.49% 的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了 99.92% 的用户
*/

func main() {
	println(longestPalindrome("babad"))
	println(longestPalindrome("cbbd"))
}
