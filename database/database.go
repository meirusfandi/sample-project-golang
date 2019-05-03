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

func GetUser(idx int) User {
	db := ConnectDB()
	data, err := db.Query("select * from pengguna where id=?", idx)
	if err != nil {
		panic(err.Error())
	}

	result := User{}
	for data.Next() {
		var id int
		var uname string
		var fullname string
		var pass string
		var email string
		err = data.Scan(&id, &uname, &pass, &email, &fullname)
		if err != nil {
			panic(err.Error())
		}
		result.Id = id
		result.Fullname = fullname
		result.Email = email
		result.Username = uname
		result.Password = pass
	}

	return result
}

func AddUser(user User) {
	db := ConnectDB()
	defer db.Close()
	rows, err := db.Prepare("INSERT INTO pengguna(id, username, password, email, fullname) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	rows.Exec(nil, user.Username, user.Password, user.Email, user.Fullname)
	defer rows.Close()
}

func UpdateUser(user User) {
	db := ConnectDB()
	defer db.Close()
	rows, err := db.Prepare("UPDATE pengguna SET password=?, email=?, fullname=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	rows.Exec(user.Password, user.Email, user.Fullname, user.Id)
	defer rows.Close()
}

func DeleteUser() {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
