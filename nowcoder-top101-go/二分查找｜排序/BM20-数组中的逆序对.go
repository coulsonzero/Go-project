package 二分查找_排序

/**
 * 题目：数组中的逆序对
 * 难度：中等
 * 描述：在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
		输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007
 * 时间复杂度：O(nlogN)，空间复杂度：O(n)
 * 输入：[1,2,3,4,5,6,7,0]
 * 输出：7
 * 输入：[1,2,3]
 * 输出：0
*/

func InversePairs(nums []int) int {
	// write code here
	const mod = 1000000007
	var count int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				count++
			}
		}
	}

	return count % mod
}
