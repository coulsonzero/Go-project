package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Printf("host: %s, post: %s \n", host, port)
}
