package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
	"time"
)

type TokenClaims struct {
	UserID string `json:"id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type LoginResponse struct {
	AccessToken  string `json : "access_token"`
	Message      string `json : "message"`
	RefreshToken string `json : "refresh_token"`
}

func main() {
	m := make(map[string]string)
	e := echo.New()

	e.GET("/signup", func(c echo.Context) error {
		id := c.QueryParam("id")
		password := c.QueryParam("password")

		if i := m[id]; i != "" {
			return c.NoContent(202)
		}
		m[id] = password
		return c.NoContent(200)
	})

	e.GET("/signin", func(c echo.Context) error {
		id := c.QueryParam("id")
		password := c.QueryParam("password")
		if i := m[id]; i == password {
			tc := TokenClaims{
				UserID: id,
				Role:   "user",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: jwt.At(time.Now().Add(time.Minute)),
				},
			}
			tcToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tc)
			accessToken, err := tcToken.SignedString([]byte("qlalfzl"))
			if err != nil {
				panic(err)
			}
			fmt.Println(accessToken)
			return c.JSON(200, "로그인에 성공하셨습니다!")
		}
		return c.NoContent(404)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
