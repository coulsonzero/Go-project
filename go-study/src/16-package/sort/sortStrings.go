package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zhangyunhao116/pdqsort"
)

func main() {
	demo()

}

func demo() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s) // [Alpha Bravo Delta Go Gopher Grin]
	fmt.Println(s)

	s2 := []string{"3", "30", "34", "5", "9"}
	// sort.Strings(s2)
	// fmt.Println(s2)
	sort.Sort(sort.Reverse(sort.StringSlice(s2))) // [9 5 34 30 3]
	fmt.Println(strings.Join(s2, ""))

	s3 := []string{"3", "30", "34", "5", "9"}
	pdqsort.Slice(s3) // [3 30 34 5 9]
	fmt.Println(s3)
}

func SortStringReverse(s []string) {
	sort.Slice(s, func(i, j int) bool { return s[i]+s[j] > s[j]+s[i] })
}
