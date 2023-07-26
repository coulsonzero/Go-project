package main

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func main() {
	println(GetUUID()) // 1443144a-45a6-433a-a854-f5dda05019f5
}
