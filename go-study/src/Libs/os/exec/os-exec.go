package main

import (
	"fmt"
	"os/exec"
)

func main() {
	goCmd3()
}

/**
 * go程序执行cmd命令
 */

func goCmd() {
	cmd := exec.Command("ls")
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

// 简写
func goCmd2() {
	res, _ := exec.Command("ls").Output()
	fmt.Println(string(res))
}

// 带参数
func goCmd3() {
	lsOut, _ := exec.Command("bash", "-c", "ls -a -l -h").Output()
	fmt.Printf("> ls -a -l -h: %s", string(lsOut))
}
