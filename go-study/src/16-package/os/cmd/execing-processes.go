package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 此方法目录会着重显示
func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		log.Fatal(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
