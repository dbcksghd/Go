package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	password := os.Getenv("PASSWORD")
	dsn := "root:" + password + "@tcp(localhost:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("what")
}
