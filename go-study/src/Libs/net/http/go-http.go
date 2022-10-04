package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! \n")
}
