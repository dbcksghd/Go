package main

import (
	"database/sql"
	"fmt"
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

	e.GET("/signin", func(c echo.Context) error {
		user := User{}
		user.id = c.QueryParam("userId")
		user.password = c.QueryParam("userPassword")
		err = db.QueryRow("SELECT * from user WHERE user_id = ? AND password = ?", user.id, user.password).
			Scan(&user.name, &user.id, &user.password)
		if err != nil {
			fmt.Println("아이디나 비밀번호가 일치하지 않습니다.")
			return c.NoContent(204)
		}
		fmt.Println(user.name)
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
