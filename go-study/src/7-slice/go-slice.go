package main

import (
	"fmt"

	"github.com/wxnacy/wgo/arrays"
)

// make: 主动扩容, 返回新的slice, 原来的slice不变
// copy: 拷贝短的slice元素到另一个长的slice中, 自动截取
// append: 自动扩容

func main() {
	// slice := make([]int, 1000, 1024)
	// slice = append(slice, []int{471: 2}...)
	// fmt.Println(cap(slice))
	// toString(slice)

	sliceAppend()

}

func toString(slice ...[]int) {
	// fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
	for _, v := range slice {
		fmt.Printf("len: %d, cap: %d, slice: %v \n", len(v), cap(v), v)
	}
}

func sliceUpdate() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9}
	copy(slice, slice2)
	fmt.Printf("slice: %v, len: %d, cap: %d \n", slice, len(slice), cap(slice))
	fmt.Printf("slice2: %v, len: %d, cap: %d \n", slice2, len(slice2), cap(slice2))
}

func sliceCap() {
	slice := []int{1, 2, 3}
	toString(slice)
	// 主动扩容
	newSlice := make([]int, len(slice), 10)
	copy(newSlice, slice)
	toString(newSlice)
	// newSlice: [1 2 3], len: 3, cap: 10

	// 自动扩容
	slice = append(slice, 4, 5)
	fmt.Printf("slice: %v, len: %d, cap: %d \n", slice, len(slice), cap(slice))
}

func sliceCopy() {
	nums := []int{1, 3, 5}
	nums2 := []int{}

	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums2), cap(nums2), nums2)
	nums2 = nums[:]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums2), cap(nums2), nums2)
}

func sliceAppend() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9}
	res := append(slice, slice2...)
	fmt.Println(res)
}

func sliceInsert(slice []int, index int, value int) []int {
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	return slice
}

func sliceInsert2(slice []int, index int, value int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func sliceRemove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func sliceReverse(s []int) []int {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}

func equal(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func sliceMerge(nums1, nums2 []int, startIndex int) []int {
	for i, _ := range nums2 {
		if startIndex+i > len(nums1)-1 {
			break
		}
		nums1[startIndex+i] = nums2[i]
	}
	return nums1
}

func contains() {
	slice := []int64{1, 3, 5, 7}
	res := arrays.ContainsInt(slice, 8)
	fmt.Println(res) // -1
}
