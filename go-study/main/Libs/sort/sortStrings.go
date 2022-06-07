package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s) // [Alpha Bravo Delta Go Gopher Grin]
	fmt.Println(s)
}
