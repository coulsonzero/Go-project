package main

import (
	"fmt"
	"strings"
)

func main() {

	res := strings.Trim("   hello world !   ", " ")
	fmt.Println(res)

	ans := strings.TrimSpace("   hello world !   ")
	fmt.Println(ans)
}

/*
   hello world !
hello world !
*/
