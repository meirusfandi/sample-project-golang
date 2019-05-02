package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// id       int    `json:"id"`
	// username string `json:"username"`
	// password string `json:"password"`
	// email    string `json:"email"`
	// name     string `json:"name"`

	id       int
	username string
	password string
	email    string
	name     string
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/golang")
	if err != nil {
		return nil
	}

	return db
}

func GetAllUser() []User {
	//get data from database
	db := ConnectDB()
	fmt.Println("connecting to database has been successfully")

	data, err := db.Query("select * from pengguna")
	if err != nil {
		fmt.Println(err.Error())
		checkErr(err)
	}
	fmt.Println("get data from database has been successfully")

	var result []User

	for data.Next() {
		var each = User{}
		var err = data.Scan(&each.id, &each.username, &each.password, &each.name, &each.email)
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

	fmt.Println(result)

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
