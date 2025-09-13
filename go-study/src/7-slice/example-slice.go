package main

import "fmt"

func main() {
	slice := make([]int, 0)
	slice = append(slice, 4) // 末尾新增一个元素
	slice = append(slice, 4, 5, 6)

	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)

	slice = insert(slice, 1, 7)
	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
}

type sl interface {
	int | int64 | float64 | string | bool
}

func insert[T sl](slice []T, i int, v T) []T {
	return append(slice[:i], append([]T{v}, slice[i:]...)...)
}
