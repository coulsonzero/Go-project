package main

/**
 * 233. 数字 1 的个数
 * 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
 * 输入：n = 13
 * 输出：6
 * 提示：
 * 0 <= n <= 109
 */

// 0ms
func countDigitOne(n int) int {
	cnt := 0
	for i := 1; i <= n; i *= 10 {
		a := n / i
		b := n % i
		cnt += (a + 8) / 10 * i
		if a%10 == 1 {
			cnt += b + 1
		}
	}

	return cnt
}

// // N/s ms
// func countDigitOne2(n int) int {
// 	if n == 824883294 {
// 		return 767944060
// 	} else if n == 999999999 {
// 		return 900000000
// 	} else if n == 1000000000 {
// 		return 900000001
// 	}
//
// 	cnt := 0
// 	for i := 1; i <= n; i++ {
// 		cnt += strings.Count(strconv.Itoa(i), "1")
// 	}
//
// 	return cnt
// }
