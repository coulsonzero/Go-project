package main

import (
	"errors"
	"fmt"
)

func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		return errors.New("网络延迟").Error()
	}
	return ""
}

func main() {
	fmt.Println(defineerr(150))
}
