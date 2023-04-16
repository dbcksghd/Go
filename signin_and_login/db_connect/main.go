package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type User struct {
	name     string
	user_id  int
	birthday string
}

func main() {
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
