package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

type User struct {
	name     string
	user_id  int
	birthday string
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
		newUser.user_id, _ = strconv.Atoi(c.QueryParam("userId"))
		newUser.birthday = c.QueryParam("userBirthday")
		_, err = db.Exec("INSERT INTO user (name, user_id, birthday) VALUES (?, ?, ?)",
			newUser.name, newUser.user_id, newUser.birthday)
		if err != nil {
			return c.NoContent(204)
		}
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
