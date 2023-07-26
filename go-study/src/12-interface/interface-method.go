package main

import (
	"fmt"
	"reflect"
	"sort"
)

// type any = interface{}

// interface 作为 不定参数
func change(t ...any) {
	Type := reflect.TypeOf(t).Kind()
	switch Type {
	case reflect.Slice:
		fmt.Println(t...)
	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("Int")
	}
}

// interface 作为 可选参数
func sortInts(nums []int, reverse ...any) {
	if len(reverse) == 1 && reverse[0] == true {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		return
	}
	sort.Ints(nums)
}

func main() {
	slice := []int{1, 5, 3}
	// change(slice)
	// sortInts(slice)			// [1, 3, 5]
	sortInts(slice, true) // [5, 3, 1]
	fmt.Println(slice)
}
