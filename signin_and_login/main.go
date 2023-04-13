package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	m := make(map[string]string)
	e := echo.New()
	e.POST("/signin", func(c echo.Context) error {
		email := c.QueryParam("email")
		pwd := c.QueryParam("password")
		m[email] = pwd
		val, isTrue := m[email]
		defer fmt.Println(val, isTrue) //잘 출력해줌 !!
		return c.NoContent(200)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
