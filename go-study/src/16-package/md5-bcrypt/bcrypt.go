// passwords.go
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 解密
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

/*
Password: secret
Hash:     $2a$14$Om4.1ottCO9NjNp6T3BaJ.FF7WgPYa6Pp/XPjHefXUnE6ZInI3VhS
Match:    true
*/
