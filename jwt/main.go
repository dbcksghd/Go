package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type TokenClaims struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func main() {
	tc := TokenClaims{
		UserID: "chanhong1206",
		Name:   "yoochanhong",
		Role:   "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
		},
	}

	tcToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tc)
	accessToken, err := tcToken.SignedString([]byte("qlalfzl"))
	if err != nil {
		panic(err)
	}
	fmt.Println(accessToken)
}
