package main

import (
	"strconv"
	"strings"
)

/**
 * 2243. 计算字符串的数字和
 * 难度：简单
 * 输入：s = "11111222223", k = 3
 * 输出："135"
 * 解释：
 * - 第一轮，将 s 分成："111"、"112"、"222" 和 "23" 。
 *   接着，计算每一组的数字和：1 + 1 + 1 = 3、1 + 1 + 2 = 4、2 + 2 + 2 = 6 和 2 + 3 = 5 。
 *   这样，s 在第一轮之后变成 "3" + "4" + "6" + "5" = "3465" 。
 * - 第二轮，将 s 分成："346" 和 "5" 。
 *   接着，计算每一组的数字和：3 + 4 + 6 = 13 、5 = 5 。
 *   这样，s 在第二轮之后变成 "13" + "5" = "135" 。
 * 现在，s.length <= k ，所以返回 "135" 作为答案。
 *
 * 输入：s = "00000000", k = 3
 * 输出："000"
 * 解释：
 * 将 "000", "000", and "00".
 * 接着，计算每一组的数字和：0 + 0 + 0 = 0 、0 + 0 + 0 = 0 和 0 + 0 = 0 。
 * s 变为 "0" + "0" + "0" = "000" ，其长度等于 k ，所以返回 "000" 。
 */

func digitSum(s string, k int) string {
	for len(s) > k {
		var build strings.Builder
		n := len(s)
		for i := 0; i < n; i += k {
			var sum int
			for j := 0; j < k && i+j < n; j++ {
				sum += int(s[i+j] - '0')
			}
			build.WriteString(strconv.Itoa(sum))
		}
		s = build.String()
	}
	return s
}

/*
class Solution:
    def digitSum(self, s: str, k: int) -> str:
        while k < len(s):
            t = ''
            for i in range(0, len(s), k):
                t += str(sum(map(int, s[i:k+i])))
            s = t
        return s
*/
