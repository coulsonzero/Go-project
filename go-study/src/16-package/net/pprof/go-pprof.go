package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	for {
		Sum()
	}
}

func Sum() {
	ans := 0
	for i := 0; i < 10000; i++ {
		ans += i
	}
	fmt.Println(ans)
}
