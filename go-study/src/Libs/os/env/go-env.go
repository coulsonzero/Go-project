package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("HOST"), os.Getenv("PORT"))
	os.Setenv("PORT", "8070")
	fmt.Println(os.Getenv("PORT"))
}

func demo() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
}
