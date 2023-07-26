package main

import "fmt"

type sl interface {
	int | int64 | float64 | string | bool
}

func SlicePrint[T sl](slice []T) {
	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
}

func SliceInsert[T sl](slice []T, index int, value T) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

func SliceRemove[T sl](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
