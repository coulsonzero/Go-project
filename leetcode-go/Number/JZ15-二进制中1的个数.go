package main

import (
	"strconv"
	"strings"
)

/**
 * 二进制中1的个数
 * 输入：n = 11 (控制台输入 00000000000000000000000000001011)
 * 输出：3
 * 十进制转二进制: strconv.FormatInt(n, 2): int64 -> string
 */

func hammingWeight(num uint32) int {
	return strings.Count(strconv.FormatInt(int64(num), 2), "1")
}

func hammingWeight2(num uint32) int {
	var res int
	for n := int(num); n != 0; n /= 2 {
		res += n % 2
	}
	return res
}
