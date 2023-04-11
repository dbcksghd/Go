package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
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
}
