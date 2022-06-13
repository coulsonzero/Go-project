package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpHandle()
}

func helloHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
	/*
		$ curl -s http://localhost:80/
		Hello, you've requested: /

	*/
}

func httpHandle() {
	fmt.Println("please open http://127.0.0.1:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("hello world! -- time: %s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("%v\n", s)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http listen: ", err)
	}
}
