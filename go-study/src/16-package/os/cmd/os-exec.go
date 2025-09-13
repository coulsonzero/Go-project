package main

import (
	"fmt"
	"os/exec"
)

func main() {
	goCmd()
}

/**
 * go程序执行cmd命令
 */

func goCmd() {
	res, _ := exec.Command("ls").Output()
	fmt.Println(string(res))
}

// 带参数
func goCmd2() {
	res, _ := exec.Command("bash", "-c", "ls -l").Output()
	fmt.Printf(string(res))
}
