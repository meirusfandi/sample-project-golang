package database

import (
	_ "github.com/go-sql-driver/mysql"
)

// type User struct {
// 	id       int    `json:"id"`
// 	username string `json:"username"`
// 	password string `json:"password"`
// 	email    string `json:"email"`
// 	name     string `json:"name"`
// }

// func ConnectDB() (*sql.DB, error) {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/golang")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

func GetAllUser() {
	//get data from database
	// db, err := ConnectDB()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// defer db.Close()

	// rows, err := db.Query("select * from pengguna")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// defer rows.Close()

	// var result []User

	// for rows.Next() {
	// 	var each = User{}
	// 	var err = rows.Scan(&each.id, &each.username, &each.password, &each.name, &each.email)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}

	// 	result = append(result, each)
	// }

	// if err = rows.Err(); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(result)

	// return result
}
