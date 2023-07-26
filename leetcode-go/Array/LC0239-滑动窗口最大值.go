package main

// k大于50000时超过时间限制
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		ans[i] = maxNum(nums[i : i+k])
	}

	return ans
}

func maxNum(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {

}
