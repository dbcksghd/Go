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
	rfm := make(map[string]string)
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
			rfm[id] = refreshToken
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

	e.GET("/checkRefreshToken", func(c echo.Context) error {
		refreshToken := c.QueryParam("refreshToken")
		token, err := jwt.ParseWithClaims(refreshToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("qlalfzl"), nil
		})
		//토큰 기간 만료 또는 에러
		if !token.Valid || err != nil {
			return c.NoContent(401)
		}

		//정의해둔 구조체에 맞게 파싱
		claims, ok := token.Claims.(*TokenClaims)
		if !ok {
			return c.NoContent(401)
		}
		id := claims.UserID
		//만약 id에 맞는 리프레시 토큰이 아니라면
		if rf := rfm[id]; rf != refreshToken {
			return c.NoContent(401)
		}
		tc := TokenClaims{
			UserID: id,
			Role:   "user",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: jwt.At(time.Now().Add(time.Minute)),
			},
		}
		tcToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tc)
		accessToken, err := tcToken.SignedString([]byte("qlalfzl"))
		return c.JSON(200, map[string]interface{}{
			"access_token": accessToken,
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
