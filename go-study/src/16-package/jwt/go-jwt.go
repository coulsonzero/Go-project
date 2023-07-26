package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

/**
 * token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{}
 * token.SignedString([]byte(jwtSecret))
 *
 * jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error)
 * claims, ok := token.Claims.(*Claims); ok && token.Valid
 */

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

const jwtSecret = "AllYourBest"

func main() {
	token, _ := GenerateToken2()
	fmt.Println(token)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJleHAiOjE2NzE5NjI4MDYsImlzcyI6Imp3dC1kZW1vIn0.fPILbgYX6SAvBDy-VOLo98YICru4s2ko70NEX2O1mxE
	parseToken, _ := ParseToken(token)
	fmt.Printf("%v\n", parseToken)
	// &{1 { 1671962806  0 jwt-demo 0 }}

}

// GenerateToken 签发Token
func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  1516239022,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))

	fmt.Println(tokenString)
	return tokenString, err
}

func GenerateToken2() (string, error) {
	claims := Claims{
		Id: "1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "jwt-demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err

}

// ParseToken 验证token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
