package main

/**
 * 输入：nums = [2,0,2,1,1,0]
 * 输出：[0,0,1,1,2,2]
 * 输入：nums = [2,0,1]
 * 输出：[0,1,2]
 */

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

// quicksort 快速排序
func quickSort(arr []int, l, r int) {
	if l < r {
		pivot := arr[r]
		j := l - 1
		for i := l; i < r; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[r], arr[j+1] = arr[j+1], arr[r]
		j += 1

		quickSort(arr, l, j-1)
		quickSort(arr, j+1, r)
	}
}
