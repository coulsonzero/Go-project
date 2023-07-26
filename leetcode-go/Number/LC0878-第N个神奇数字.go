package main

/**
 * question: 878. 第 N 个神奇数字
 * 难度：困难
 * descript: 一个正整数如果能被 a 或 b 整除，那么它是神奇的。
 *           给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 109 + 7 取模 后的值
 * input: n = 1, a = 2, b = 3
 * output: 2
 * input: n = 4, a = 2, b = 3
 * output: 6
 * 解释：符合条件的为[2,3,4,6] 第4个为 6
 */

// 二分查找
func nthMagicalNumber(n, a, b int) int {
	const mod int = 1e9 + 7
	l, r := 2, n*max(a, b)

	for l < r {
		mid := l + (r-l)/2
		cnt := mid/a + mid/b - mid/lcm(a, b)
		if cnt < n {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 超出时间限制
func nthMagicalNumber2(n int, a int, b int) int {
	const mod int = 1e9 + 7
	res := make([]int, 0)

	for i := 2; i > 1; i++ {
		if i%a == 0 || i%b == 0 {
			res = append(res, i)
		}

		if len(res) == n {
			break
		}
	}
	// fmt.Println(res)
	return res[len(res)-1] % mod
}

func main() {
	println(nthMagicalNumber(1, 2, 3))                // 2
	println(nthMagicalNumber(4, 2, 4))                // 8
	println(nthMagicalNumber(7, 2, 4))                // 14
	println(nthMagicalNumber(100, 5, 7))              // 320
	println(nthMagicalNumber(10000000, 40000, 40000)) // 可能会超出时间限制, 999997207

}
