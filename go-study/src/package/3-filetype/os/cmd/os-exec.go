package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdTail()
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

func goCmd4() {
	res, _ := exec.Command("git", "status").Output()
	fmt.Println(string(res))
}

func cmdTail() {
	res, _ := exec.Command("tail", "-n", "+18", "bill.csv").Output()
	fmt.Println(string(res))
}

func Exec() {
	out := exec.Command("bash", "-c", "tail -n +18 bill.csv >> res.csv")
	_ = out.Start()
	_ = out.Wait()
}
