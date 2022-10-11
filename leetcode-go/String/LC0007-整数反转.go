package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverse2(-123)) // -321
	fmt.Println(reverse2(123))  // 321
	fmt.Println(reverse2(120))  // 21
	fmt.Println(reverse2(0))    // 0

}

// 解法一
func reverse2(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res%10 != x%10 {
			return 0
		}
		x /= 10
	}
	return res
}

// 解法二
func reverse(x int) int {
	ReverseString := func(s []byte) string {
		j := len(s) - 1
		for i := 0; i < j; i++ {
			s[i], s[j] = s[j], s[i]
			j--
		}
		return string(s)
	}

	if strings.HasPrefix(strconv.Itoa(x), "-") {
		s := strconv.Itoa(int(math.Abs(float64(x))))
		res, _ := strconv.Atoi("-" + ReverseString([]byte(s)))
		return res
	}
	ret, _ := strconv.Atoi(ReverseString([]byte(strconv.Itoa(x))))
	return ret
}

/*
func ReverseString(s []byte) string {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return string(s)
}
*/

/*
// python3

class Solution:
    def reverse(self, x: int) -> int:
        num = int(str(abs(x))[::-1])
        s = -num if x<0 else num
        return s if -2**31 <= s <= 2**31-1 else 0
*/
