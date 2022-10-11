package main

import (
	"fmt"
	"regexp"
)

const str string = `{"confirmed":"4234","died":"9","crued":"4179"}`

func main() {
	// reg := regexp.MustCompile(`"confirmed":"(\d+)","died":"(\d+)","crued":"(\d+)"`)
	reg := regexp.MustCompile(`"[\w]+":"[\d]+"`)

	// res := reg.FindString(str)			  // "confirmed":"4234"
	// res := reg.FindStringSubmatch(str)     // ["confirmed":"4234"]
	res := reg.FindAllStringSubmatch(str, -1) // [["confirmed":"4234"] ["died":"9"] ["crued":"4179"]]

	fmt.Println(res)
}
