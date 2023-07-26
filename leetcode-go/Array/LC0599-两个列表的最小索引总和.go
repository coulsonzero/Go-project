package main

import (
	"math"
	"sort"
)

// 执行时间：52ms，击败 5.61% 的用户; 内存占用：6.9MB，击败 82.24% 的用户
func findRestaurant(list1 []string, list2 []string) []string {
	indexOf := func(arr []string, value string) int {
		for i, v := range arr {
			if v == value {
				return i
			}
		}
		return -1
	}

	m := make(map[string]int)
	for i, v := range list1 {
		k := indexOf(list2, v)
		if k != -1 {
			// map统计索引和
			m[v] = k + i
		}
	}

	// fmt.Println(m)

	// 根据map中的value排序，输出最小的value对应的key(可能存在多个)
	var vArr []int
	for _, v := range m {
		vArr = append(vArr, v)
	}
	sort.Ints(vArr)

	var kArr []string
	for k, v := range m {
		if v == vArr[0] {
			kArr = append(kArr, k)
		}
	}
	return kArr
}

// 执行时间：20ms，击败 99.7% 的用户; 内存占用：6.8MB，击败 90.65% 的用户
func findRestaurant2(list1, list2 []string) (ans []string) {
	index := make(map[string]int, len(list1))
	for i, s := range list1 {
		index[s] = i
	}

	indexSum := math.MaxInt32
	for i, s := range list2 {
		if j, ok := index[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{s}
			} else if i+j == indexSum {
				ans = append(ans, s)
			}
		}
	}
	return
}
