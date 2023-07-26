package main

import (
	"fmt"
	"regexp"
)

/**
 * reg := regexp.MustCompile(s string)
 *
 * reg.FindString(str) string
 * reg.FindAllString(str, -1) []string
 * reg.FindAllStringSubmatch(str, -1) [][]string
 */

const str string = `{"confirmed":"4234","died":"9","crued":"4179"}`

func main() {
	// reg := regexp.MustCompile(`"confirmed":"(\d+)","died":"(\d+)","crued":"(\d+)"`)
	reg := regexp.MustCompile(`"[\w]+":"[\d]+"`)

	// res := reg.FindString(str)			  // "confirmed":"4234"
	// res := reg.FindStringSubmatch(str)     // ["confirmed":"4234"]
	res := reg.FindAllStringSubmatch(str, -1) // [["confirmed":"4234"] ["died":"9"] ["crued":"4179"]]

	fmt.Println(res)
}

func demo() {
	word := "a123bc34d8ef34"
	reg := regexp.MustCompile(`\d+`)
	res := reg.FindAllString(word, -1)          // ["123", "34", "8", "34"]
	res2 := reg.FindAllStringSubmatch(word, -1) // [["123"], ["34"], ["8"], ["34"]]
	res3 := reg.FindString(word)                // "123"
	fmt.Println(res, res2, res3)
}
