package 动态规划

func minCostClimbingStairs(cost []int) int {
	// write code here
	fast, slow := 0, 0
	for i := 2; i <= len(cost); i++ {
		fast, slow = slow, min(slow+cost[i-1], fast+cost[i-2])
	}
	return slow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
