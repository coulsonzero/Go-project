package main

/**
 * question: 1775. 通过最少操作次数使数组的和相等
 * description: 给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）
 * 				每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）
 * 				请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1
 *
 * 输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
 * 输出：3
 * 解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
 * - 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
 * - 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
 * - 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。
 *
 * 输入：nums1 = [1,1,1,1,1,1,1], nums2 = [6]
 * 输出：-1
 * 解释：没有办法减少 nums1 的和或者增加 nums2 的和使二者相等。
 */

func minOperations(nums1 []int, nums2 []int) int {
	sum := 0
	ans := 0
	cnt := make([]int, 6)
	for _, v := range nums1 {
		cnt[v-1]++
		sum += v
	}
	for _, v := range nums2 {
		cnt[6-v]++
		sum -= v
	}
	if sum == 0 {
		return 0
	} else if sum < 0 {
		return minOperations(nums2, nums1)
	}

	for i := 5; i >= 0; i-- {
		if sum <= cnt[i]*i {
			return ans + (sum+i-1)/i
		}
		sum -= cnt[i] * i
		ans += cnt[i]
	}

	return -1
}
