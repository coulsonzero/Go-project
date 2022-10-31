package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// s := "A3B2" // AAABB
	// s := "{A3B1{C}3}3" // AAABCCCAAABCCCAAABCCC
	// s := "{{A3B1{C}3}3}2" // AAABCCCAAABCCCAAABCCCAAABCCCAAABCCCAAABCCC
	s := "{{{A3B1{C}3}3}2D3}2" // AAABCCCAAABCCCAAABCCCAAABCCCAAABCCCAAABCCCDDDAAABCCCAAABCCCAAABCCCAAABCCCAAABCCCAAABCCCDDD

	var temp string
	var c string
	for f := 0; f < len(s); f++ {
		if unicode.IsLetter(rune(s[f])) {
			c = string(s[f])
		} else if unicode.IsDigit(rune(s[f])) {
			if c != "" {
				num, _ := strconv.Atoi(string(s[f]))
				temp += strings.Repeat(c, num)
				c = ""
			} else {
				num, _ := strconv.Atoi(string(s[f]))
				temp = strings.Repeat(temp, num)
			}
		}
	}

	fmt.Println(temp)
}
