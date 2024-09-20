// Authentication
// Berikut ini implementasi sederhana dari authentication berbasis JWT di Go menggunakan package github.com/dgrijalva/jwt-go:

package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// Define a struct to represent claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate JWT token
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Function to validate JWT token
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func main() {
	token, _ := GenerateJWT("miftahariss")
	fmt.Println("Generated Token:", token)

	// Validate the token
	claims, err := ValidateJWT(token)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	fmt.Println("Authenticated User:", claims.Username)
}
