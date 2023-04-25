package main

import "github.com/labstack/echo/v4"

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
	e.Logger.Fatal(e.Start(":8080"))
}
