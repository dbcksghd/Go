package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"os"
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

	newUser := User{user_id: 1265, name: "강민", birthday: "20060629"}
	_, err = db.Exec("INSERT INTO user (name, user_id, birthday) VALUES (?, ?, ?)",
		newUser.name, newUser.user_id, newUser.birthday)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
