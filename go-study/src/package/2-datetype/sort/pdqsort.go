package main

import (
	"fmt"
	"github.com/zhangyunhao116/pdqsort"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	nums := []int{3, 1, 2, 4, 5, 9, 8, 7}

	pdqsort.Slice(nums)
	fmt.Printf("sort_reslut = %v\n", nums)

	searchResult := pdqsort.Search(nums, 5)
	fmt.Println(searchResult)

	isSort := pdqsort.SliceIsSorted(nums)
	fmt.Println(isSort)

}
