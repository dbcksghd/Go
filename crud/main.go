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
	//db 연결
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//CREATE
	newUser := User{UserId: 1235, Name: "길근우", Birthday: "20060412"}
	_, err = db.Exec("INSERT INTO user (user_id, name, birthday) VALUES (?, ?, ?)", newUser.UserId, newUser.Name, newUser.Birthday)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(newUser.UserId, newUser.Name, newUser.Birthday) // 1235 길근우 20060412

	//READ
	rows, err := db.Query("SELECT *FROM user")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var users []User //유저 리스트 생성
	for rows.Next() {
		var user User                                                              //유저 객체
		if err = rows.Scan(&user.UserId, &user.Name, &user.Birthday); err != nil { //만약 받아온 값에 에러가 있다면
			log.Fatalln(err) //로그 찍어주기
		}
		users = append(users, user) //없으면 유저 리스트에 유저 추가
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	for _, user := range users { //유저 리스트 돌면서 가져온 값 다 출력해주기
		fmt.Println(user.UserId, user.Name, user.Birthday)
	}

	//UPDATE
	updateUser := User{
		UserId:   1235,
		Name:     "이정호",
		Birthday: "20061209",
	} //업뎃할 유저
	_, err = db.Exec("UPDATE user SET name = ?, birthday = ? WHERE user_id=?",
		updateUser.Name, updateUser.Birthday, updateUser.UserId)
	if err != nil {
		log.Fatalln(err)
	}

	//DELETE
	deleteUser := User{
		UserId:   1234,
		Name:     "유찬홍",
		Birthday: "20061206",
	}
	_, err = db.Exec("DELETE FROM user WHERE user_id=?", deleteUser.UserId)
	if err != nil {
		log.Fatalln(err)
	}
}
