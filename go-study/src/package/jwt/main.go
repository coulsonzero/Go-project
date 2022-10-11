package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func main() {
	// GenerateToken()
	fmt.Println(ParseToken())
	// GenerateToken2()
}

// GenerateToken 签发Token
func GenerateToken() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  1516239022,
	})

	// Sign and get the complete encoded token as a string using the secret
	mySigningKey := "AllYourBest"

	tokenString, _ := token.SignedString([]byte(mySigningKey))

	fmt.Println(tokenString)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSIsInN1YiI6IjEyMzQ1Njc4OTAifQ.Zifx_niuGRaw01fZLbvcjSfuCP48R7G19G3EJ3Rcr7g
}

func GenerateToken2() {
	claims := Claims{
		Id: "1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "jwt-demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := "AllYourBest"
	tokenString, _ := token.SignedString([]byte(mySigningKey))

	fmt.Println(tokenString)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjY0NjA2MTYxLCJpc3MiOiJqd3QtZGVtbyJ9.FZuv42p5b_bf6YFXjoeN-gfuhqEDOGvvfg26FgxOwzA
}

// ParseToken 验证token
func ParseToken() (*Claims, error) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSIsInN1YiI6IjEyMzQ1Njc4OTAifQ.Zifx_niuGRaw01fZLbvcjSfuCP48R7G19G3EJ3Rcr7g"

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBast"), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
