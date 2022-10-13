package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	"log"
)

func DemoCmdOutput() {
	output, err := pro.CmdOutput("ls")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func main() {
	encode := pro.MD5Encode("admin123")
	fmt.Println(encode)
}
