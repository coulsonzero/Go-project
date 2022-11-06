package main

import (
	"strconv"
	"strings"
)

/**
 * 165. 比较版本号
 * 难度：中等
 * 返回规则如下：
 * 如果 version1 > version2 返回 1，
 * 如果 version1 < version2 返回 -1，
 * 除此之外返回 0。
 */

func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

func main() {
	println(compareVersion("1.0.1", "1"))     // 1
	println(compareVersion("0.1", "1.1"))     // -1
	println(compareVersion("1.0.1", "1.001")) // 0
}
