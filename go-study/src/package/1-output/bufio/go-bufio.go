package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// 模拟shell脚本命令

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if text == "ls" {
			execCmd()
		}

		if text == "exit" {
			_ = os.Stdin.Close()
		} else {
			fmt.Print(">>> ")
		}

	}
	fmt.Println("========= close ==========")
}

func execCmd() {
	lsOut, _ := exec.Command("bash", "-c", "ls -a -l -h").Output()
	fmt.Printf("> ls -a -l -h: %s", string(lsOut))
}
