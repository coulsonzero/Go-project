package main

import "strings"

/**
* 1704. 判断字符串的两半是否相似
* 难度：简单
* 给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b
* 两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。
* 如果 a 和 b 相似，返回 true ；否则，返回 false
*
* 输入：s = "book"
* 输出：true
* 解释：a = "bo" 且 b = "ok" 。a 中有 1 个元音，b 也有 1 个元音。所以，a 和 b 相似。
* 输入：s = "textbook"
* 输出：false
* 解释：a = "text" 且 b = "book" 。a 中有 1 个元音，b 中有 2 个元音。因此，a 和 b 不相似。注意，元音 o 在 b 中出现两次，记为 2 个。
 */

func halvesAreAlike(s string) bool {
	temp := "aeiouAEIOU"
	cntl, cntr := 0, 0
	l, r := 0, len(s)-1
	for l < r {
		if strings.Contains(temp, string(s[l])) {
			cntl++
		}
		if strings.Contains(temp, string(s[r])) {
			cntr++
		}
		l++
		r--
	}
	return cntl == cntr
}
