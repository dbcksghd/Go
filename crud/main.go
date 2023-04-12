package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type User struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func main() {
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	newUser := User{UserId: 1235, Name: "길근우", Birthday: "20060412"}
	_, err = db.Exec("INSERT INTO user (user_id, name, birthday) VALUES (?, ?, ?)", newUser.UserId, newUser.Name, newUser.Birthday)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(newUser.UserId, newUser.Name, newUser.Birthday) // 1235 길근우 20060412
}
