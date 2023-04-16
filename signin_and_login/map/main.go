package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	m := make(map[string]string)
	e := echo.New()
	e.POST("/signup", func(c echo.Context) error {
		email := c.QueryParam("email")
		pwd := c.QueryParam("password")
		if i := m[email]; i != "" {
			return c.JSON(202, "이미 있음")
		}
		m[email] = pwd
		val, isTrue := m[email]
		defer fmt.Println(val, isTrue) //잘 출력해줌 !!
		return c.NoContent(201)
	})
	e.POST("/signin", func(c echo.Context) error {
		email := c.QueryParam("email")
		pwd := c.QueryParam("password")
		if i := m[email]; i == pwd {
			return c.JSON(200, "로그인에 성공하셨습니다!")
		}
		return c.JSON(404, "실패")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
