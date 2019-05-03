package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	email    string `json:"email"`
	fullname string `json:"fullname"`
}

func ConnectDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetAllUser() []User {
	//get data from database
	db := ConnectDB()
	data, err := db.Query("select * from pengguna ORDER BY id DESC")
	if err != nil {
		checkErr(err)
	}

	var result []User

	for data.Next() {
		var each = User{}
		var err = data.Scan(&each.id, &each.username, &each.password, &each.email, &each.fullname)
		if err != nil {
			fmt.Println(err.Error())
			checkErr(err)
		}

		fmt.Println("username : " + each.username)
		fmt.Println("password : " + each.password)
		fmt.Println("name : " + each.fullname)
		fmt.Println("email : " + each.email)

		result = append(result, each)
	}

	if err = data.Err(); err != nil {
		fmt.Println(err.Error())
		checkErr(err)
	}

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
