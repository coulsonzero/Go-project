package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Printf("nums: %v \n", arr)
}
