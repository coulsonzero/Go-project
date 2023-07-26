package main

func productExceptSelf(nums []int) []int {
	// 非零乘积
	p := 1
	// 非零个数
	cnt := 0
	// 零的个数
	zeroCnt := 0
	for _, v := range nums {
		if v != 0 {
			p *= v
			cnt++
		} else {
			zeroCnt++
		}
	}

	if cnt == 0 || zeroCnt > 1 {
		p = 0
	}

	// 1个0其他为非零的乘积，2个零及以上结果均为零
	var ans []int
	for _, v := range nums {
		if zeroCnt > 0 {
			if v == 0 {
				ans = append(ans, p)
			} else {
				ans = append(ans, 0)
			}
		} else {
			ans = append(ans, p/v)
		}

	}
	return ans
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	l, r := 1, 1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	for i := 0; i < n; i++ {
		ans[i] *= l
		l *= nums[i]

		ans[n-1-i] *= r
		r *= nums[n-1-i]
	}

	return ans
}
