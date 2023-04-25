package main

import (
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
	AccessToken  string `json:"access_token"`
	Message      string `json:"message"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	m := make(map[string]string)
	//rfm := make(map[string]string)
	e := echo.New()

	e.POST("/signup", func(c echo.Context) error {
		id := c.QueryParam("id")
		password := c.QueryParam("password")

		if i := m[id]; i != "" {
			return c.NoContent(202)
		}
		m[id] = password
		return c.NoContent(200)
	})

	e.POST("/signin", func(c echo.Context) error {
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
			rf := TokenClaims{
				UserID: id,
				Role:   "user",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
				},
			}
			rfToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &rf)
			refreshToken, err := rfToken.SignedString([]byte("qlalfzl"))
			if err != nil {
				panic(err)
			}
			loginResponse := LoginResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				Message:      "로그인에 성공하셨습니다!",
			}
			return c.JSON(200, loginResponse)
		}
		return c.NoContent(404)
	})

	e.GET("/checkToken", func(c echo.Context) error {
		accessToken := c.QueryParam("accessToken")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("qlalfzl"), nil
		})
		if !token.Valid || err != nil {
			return c.NoContent(401)
		}
		return c.NoContent(200)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
