package main

func subarrayLCM(nums []int, k int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		t := nums[i]
		for j := i; j < n; j++ {
			t = t * nums[j] / gcd(t, nums[j])
			if t == k {
				res++
			} else if t > k {
				break
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
