package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	email    string `json:"email"`
	fullname string `json:"fullname"`
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

//index for get all user on one page
func IndexUserPage(w http.ResponseWriter, r *http.Request) {
	//get data from database
	// var result = database.GetAllUser()

	db, err := ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	fmt.Println("connect db is successfully")

	rows, err := db.Query("SELECT * FROM pengguna")
	if err != nil {
		fmt.Println("select data is failed")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	fmt.Println("select data is successfully")

	// var result []User

	// for rows.Next() {
	// 	var each = User{}
	// 	var err = rows.Scan(&each.id, &each.username, &each.password, &each.fullname, &each.email)
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

	// file := path.Join("view", "index.html")
	// var templates, erro = template.ParseFiles(file)

	// if erro != nil {
	// 	http.Error(w, erro.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// erro = templates.Execute(w, result)
	// if erro != nil {
	// 	http.Error(w, erro.Error(), http.StatusInternalServerError)
	// }
}

func IndexOneUserPage(w http.ResponseWriter, r *http.Request) {
	file := path.Join("view", "user.html")
	var templates, err = template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUserPage(w http.ResponseWriter, r *http.Request) {
	file := path.Join("view", "add.html")
	var templates, err = template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUserPage(w http.ResponseWriter, r *http.Request) {
	file := path.Join("view", "update.html")
	var templates, err = template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to delete page")
}
