package Summary

import (
	"fmt"
)

func quicksort(arr []int, l, r int) {
	if r < l {
		return
	}

	pivot := partition(arr, l, r)
	quicksort(arr, l, pivot-1)
	quicksort(arr, pivot+1, r)
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	for l < r {
		for l < r && pivot < arr[r] {
			r--
		}

		for l < r && pivot > arr[l] {
			l++
		}

		arr[l], arr[r] = arr[r], arr[l]
	}

	return l
}

func main() {
	// nums := []int{10, 5, 8, 1, 4, 2}
	nums := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	// quicksort(nums, 0, len(nums)-1)
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("%v", nums) // [1 2 4 5 8 10]

}

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
