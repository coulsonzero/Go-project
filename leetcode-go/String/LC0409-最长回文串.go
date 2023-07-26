package main

import (
	"fmt"
)

/**
 * 回文串中，只有最多一个字符出现了奇数次，其余的字符都出现偶数次。
 * 将每个字符使用偶数次，使得它们根据回文中心对称。在这之后，如果有剩余的字符，我们可以再取出一个，作为回文中心。
 */

// hashmap统计次数，偶数叠加，如果出现第一个为奇数的字符后+1，之后不再增加
func longestPalindrome(s string) int {
	hashmap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}

	// 偶数数量字符不变，奇数字符数量-1
	var count int
	for _, v := range hashmap {
		count += v / 2 * 2
		// 发现了第一个出现次数为奇数的字符后，我们将 ans 增加 1, 之后不再增加
		if v%2 == 1 && count%2 == 0 {
			count++
		}

	}

	return count
}

// hashmap 统计奇数字符的个数
func longestPalindrome(s string) int {
	hashmap := make(map[rune]int)
	for _, v := range s {
		hashmap[v]++
	}

	var count int
	for _, v := range hashmap {
		count += v % 2
	}

	if count == 0 {
		return len(s)
	}
	return len(s) - count + 1
}

func main() {
	fmt.Println(longestPalindrome("abccccdd")) // 7
	fmt.Println(longestPalindrome("bananas"))  // 5
}
