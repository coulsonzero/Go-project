package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	minN := nums[0]
	maxN := nums[n-1]

	l := 0
	r := n - 1
	start := 0
	end := n - 1
	for l <= r {
		if nums[l] <= minN {
			start = l
		}
		if nums[r] >= maxN {
			end = r
		}
	}

	// fmt.Printf("%v\n", nums[start:end+1])

	return end - start
}

func main() {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
