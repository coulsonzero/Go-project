package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Printf("host: %s, post: %s \n", host, port)
}
