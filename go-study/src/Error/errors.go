package main

import (
	"errors"
	"fmt"
)

func main() {

}

func test(x int) (int, error) {
	if x < 0 {
		return -1, errors.New("can't work with 42")
	}
	return x + 2, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
