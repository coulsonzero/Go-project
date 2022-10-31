package main

import "fmt"

func main() {
	nums := []int{2, 4, 1, 6, 3}

	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(arr []int, l int, r int) {
	temp := arr[0]
	if l < r {
		arr[r], temp = temp, arr[r]
		k := partition(arr, l, r)
		quickSort(arr, l, k-1)
		quickSort(arr, k+1, r)
	}
}

func partition(arr []int, l int, r int) int {
	m := l - 1
	for i := range arr {
		if arr[i] < arr[r] {
			m += 1
			arr[m], arr[i] = arr[i], arr[m]
			arr[m+1], arr[r] = arr[r], arr[m+1]

		}
	}
	return m + 1
}
