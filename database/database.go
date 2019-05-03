package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
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
	data, err := db.Query("select * from pengguna")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	var result = []User{}

	for data.Next() {
		var each = User{}
		var err = data.Scan(&each.Id, &each.Username, &each.Password, &each.Email, &each.Fullname)
		if err != nil {
			fmt.Println(err.Error())
			checkErr(err)
		}

		result = append(result, each)
	}

	if err = data.Err(); err != nil {
		fmt.Println(err.Error())
		checkErr(err)
	}
	defer data.Close()
	return result
}

// func GetUser() User{

// }

func AddUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
