package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

type TokenClaims struct {
	UserID string   `json:"id"`
	Name   string   `json:"name"`
	Role   []string `json:"role"`
	jwt.StandardClaims
}

func main() {
	fmt.Println("what")
}
