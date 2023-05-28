package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	Id       int
	Name     string
	Birthday string
}

func main() {
	password := os.Getenv("PASSWORD")
	dsn := "root:" + password + "@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	user := User{Id: 1206, Birthday: "20061206", Name: "유찬홍"}
	db.Table("test.user").Create(&user)
	fmt.Println("what")
}
