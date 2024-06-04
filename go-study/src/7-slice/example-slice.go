package main

import "fmt"

func main() {
	slice := make([]int, 0)
	slice = append(slice, 4) // 末尾新增一个元素
	slice = append(slice, 4, 5, 6)

	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)

	// slice = SliceInsert(slice, 1, 7)
	slice = insert(slice, 1, 7)
	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
}

func SliceInsert(slice []int, index int, value int) []int {
	return append(slice[:index], append([]int{value}, slice[index:]...)...)
}

type sl interface {
	int | int64 | float64 | string | bool
}

func insert[T sl](slice []T, i int, v T) []T {
	return append(slice[:i], append([]T{v}, slice[i:]...)...)
}

func sliceInit() {
	// var slice []int
	// slice := make([]int, 0)
	slice := []int{}

	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
	// len: 0, cap: 0, slice: []
}
