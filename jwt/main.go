package main

import (
	"errors"
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

	//복호화
	claims := TokenClaims{}
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("잘못됨")
		}
		return []byte("qlalfzl"), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &claims, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t\n%#v", token.Valid, claims)
}
