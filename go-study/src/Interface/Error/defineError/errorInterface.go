package main

import "fmt"

type error interface {
	Error() string
}

type errors struct {
}

func (e errors) Error() string {
	return "网络延迟"
}

func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		return errors{}.Error()
	}
	return ""
}

func main() {
	fmt.Println(defineerr(150))
}
