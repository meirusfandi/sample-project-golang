package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func IndexUserPage(w http.ResponseWriter, r *http.Request) {

	// //get data from database
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
	// 	var err = rows.Scan(&each.ID, &each.Username, &each.Password, &each.Name, &each.Email)
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

	file := path.Join("view", "index.html")
	var templates, erro = template.ParseFiles(file)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	erro = templates.Execute(w, nil)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
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
