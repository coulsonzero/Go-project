package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("please open on http://127.0.0.1:8080")
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http listen: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	// s := "hello world"
	s := fmt.Sprintf("hello world! -- time: %s", time.Now().String())
	w.Write([]byte(s))
	log.Printf(s)
}

// http计时器中间件
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handle
		next.ServeHTTP(w, r)

		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}

func demo() {
	fmt.Println("please open on http://127.0.0.1:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		s := fmt.Sprintf("hello world! -- time: %s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		timeElapsed := time.Since(timeStart)
		log.Printf("%10v  %s\n", timeElapsed, s)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http listen: ", err)
	}
}
