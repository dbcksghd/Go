package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var (
		name     string
		birthday string
	)
	var userId int
	err = db.QueryRow("SELECT * FROM user").Scan(&userId, &name, &birthday)
	if err != nil {
		panic(err)
	}
	fmt.Println(name, userId, birthday)
}
