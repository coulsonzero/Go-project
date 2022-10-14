package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 获取终端输入
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(">>> ")
	// 持续监听
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(1)
		}
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
		fmt.Print(">>> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)

		os.Exit(1)
	}
}
