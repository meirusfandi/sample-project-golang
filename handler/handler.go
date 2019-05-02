package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func IndexUserPage(w http.ResponseWriter, r *http.Request) {
	file := path.Join("view", "index.html")
	var templates, err = template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get data from database
	db, err := ConnectDB()

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from pengguna")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer rows.Close()

	var result []User

	for rows.Next() {
		var each = User{}
		var err = rows.Scan(&each.ID, &each.Username, &each.Password, &each.Name, &each.Email)
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		result = append(result, each)
	}

	err = templates.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
