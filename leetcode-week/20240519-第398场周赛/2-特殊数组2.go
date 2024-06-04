package main

func isArraySpecial2(nums []int, queries [][]int) (ans []bool) {
	sum := make([]int, len(nums))
	for i, v := range nums[:len(nums)-1] {
		sum[i+1] = sum[i]
		if v % 2 == nums[i+1] % 2 {
			sum[i+1]++
		}
	}

	for _, q := range queries {
		l, r := q[0], q[1]
		ans = append(ans, sum[r] == sum[l])
	}
	return
}


