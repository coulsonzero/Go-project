package main

import (
	"fmt"
	"github.com/zhangyunhao116/pdqsort"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	var arr []string
	for _, num := range nums {
		arr = append(arr, strconv.Itoa(num))
	}
	fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] >= arr[j]+arr[i]
	})
	fmt.Println(strings.Join(arr, ""))

}

func demo() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s) // [Alpha Bravo Delta Go Gopher Grin]
	fmt.Println(s)

	s2 := []string{"3", "30", "34", "5", "9"}
	// sort.Strings(s2)
	// fmt.Println(s2)
	sort.Sort(sort.Reverse(sort.StringSlice(s2)))
	fmt.Println(s2)
	fmt.Println(strings.Join(s2, ""))

	s3 := []string{"3", "30", "34", "5", "9"}
	pdqsort.Slice(s3)
	fmt.Println(s3)
}

func SortStringReverse2(s []string) {
	sort.Slice(s, func(i, j int) bool {
		return s[i]+s[j] >= s[j]+s[i]
	})
}
