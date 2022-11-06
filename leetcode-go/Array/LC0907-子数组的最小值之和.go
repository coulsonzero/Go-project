package main

/**
 * Question: 907. 子数组的最小值之和
 * Description: 给定一个整数数组 arr，找到 min(b)的总和，其中 b 的范围为 arr 的每个（连续）子数组。
 * 由于答案可能很大，因此 返回答案模 10^9 + 7
 * 输入：arr = [3,1,2,4]
 * 输出：17
 * 解释：
 * 子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
 * 最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
 */

/********************************************
 * 子数组
var res [][]int
for i := 0; i < len(arr); i++ {
	for j := i; j < len(arr); j++ {
		res = append(res, arr[i:j+1])
	}
}
 ********************************************/

func sumSubarrayMins(arr []int) int {
	n, res := len(arr), 0
	// 由于答案可能很大，因此 返回答案模 10^9 + 7
	const mod = 1000000007
	for i := 0; i < n; i++ {
		l, r := i-1, i+1
		for l >= 0 && arr[l] >= arr[i] {
			l--
		}
		for r < n && arr[r] > arr[i] {
			r++
		}
		res += (i - l) * (r - i) * arr[i]
	}
	return res % mod
}

/**
 * 执行结果：通过
 * 执行用时：788 ms, 在所有 Go 提交中击败了 5.03% 的用户
 * 内存消耗：7 MB, 在所有 Go 提交中击败了 89.35% 的用户
 */

func main() {
	arr := []int{3, 1, 2, 4}
	println(sumSubarrayMins(arr))
}

// // 超出时间限制
// func sumSubarrayMins2(arr []int) int {
// 	var mod = 1000000007
// 	var sum int
// 	var res [][]int
// 	for i := 0; i < len(arr); i++ {
// 		for j := i; j < len(arr); j++ {
// 			res = append(res, arr[i:j+1])
// 		}
// 	}
// 	for i, _ := range res {
// 		sum += minNum(res[i])
// 	}
//
// 	return sum % mod
// }
//
// func minNum(nums []int) int {
// 	res := nums[0]
// 	for _, v := range nums {
// 		if v <= res {
// 			res = v
// 		}
// 	}
// 	return res
// }
