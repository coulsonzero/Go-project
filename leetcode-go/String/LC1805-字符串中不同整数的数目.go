package main

import (
	"regexp"
)

func numDifferentIntegers(word string) int {
	reg := regexp.MustCompile(`\d+`)
	arr := reg.FindAllString(word, -1)

	m := make(map[string]int)
	for _, v := range arr {
		v = removeZero(v)
		m[v]++
	}
	// fmt.Printf("%#-v", m)

	return len(m)
}

func removeZero(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != '0' || i == len(s)-1 {
			return s[i:]
		}
	}
	return s
}

func main() {
	println(numDifferentIntegers("a123bc34d8ef34"))  // 3
	println(numDifferentIntegers("leet1234code234")) // 2
	println(numDifferentIntegers("a1b01c001"))       // 1
	println(numDifferentIntegers("gi851a851q8510v")) // 2

}
