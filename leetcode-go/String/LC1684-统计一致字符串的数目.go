package main

import (
	"strings"
)

func countConsistentStrings(allowed string, words []string) int {
	ans := len(words)
	for _, w := range words {
		for _, c := range w {
			if !strings.Contains(allowed, string(c)) {
				ans--
				break
			}
		}
	}
	return ans
}

/*
func countConsistentStrings(allowed string, words []string) int {
	var cnt int
	for _, v := range words {
		v = RemoveDuplicate(v)
		if contains(allowed, v) {
			// fmt.Printf("allowed: %s, v: %s \n", allowed, v)
			cnt++
		}
	}
	return cnt
}

func RemoveDuplicate(s string) string {
	temp := strings.Split(s, "")
	sort.Strings(temp)
	var res strings.Builder
	for _, v := range temp {
		if !strings.Contains(res.String(), v) {
			res.WriteString(v)
		}
	}
	return res.String()
}

func contains(s string, sub string) bool {
	for _, v := range sub {
		if !strings.Contains(s, string(v)) {
			return false
		}
	}
	return true
}
*/

func main() {
	println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))      // 2
	println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"})) // 7

}

/*
python3
class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        return sum([len(set(i) - set(allowed)) == 0 for i in words])

javascript
var countConsistentStrings = function(allowed, words) {
    return words.filter(item => item.split("").every(ch => allowed.includes(ch))).length;
};
*/
