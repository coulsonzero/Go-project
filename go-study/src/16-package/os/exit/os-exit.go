package main

import (
	"fmt"
	"os"
)

// os.Exit(code int): 状态码0表示成功，非0表示出错。程序会立刻终止，并且 defer 的函数不会被执行。

func exit_defer() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called.
	defer fmt.Println("!")

	os.Exit(1)
}

func exit_0() {
	fmt.Println("======== start0 ========")
	os.Exit(0)
	fmt.Println("========= end0 =========")
}

/*
======== start0 ========
*/

func exit_1() {
	fmt.Println("======== start1 ========")
	os.Exit(1)
	fmt.Println("========= end1 =========")
}

/*
======== start1 ========
*/
