package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"os"
)

type User struct {
	name     string
	id       string
	password string
}

func main() {
	e := echo.New()
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e.POST("/signup", func(c echo.Context) error {
		newUser := User{}
		newUser.name = c.QueryParam("userName")
		newUser.id = c.QueryParam("userId")
		newUser.password = c.QueryParam("userPassword")
		_, err = db.Exec("INSERT INTO user (user_name, user_id, password) VALUES (?, ?, ?)",
			newUser.name, newUser.id, newUser.password)
		if err != nil {
			return c.NoContent(204)
		}
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
