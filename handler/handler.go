package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"../database"
)

//index for get all user on one page
func IndexUserPage(w http.ResponseWriter, r *http.Request) {
	//get data from database
	var result = database.GetAllUser()

	file := path.Join("view", "index.html")
	var templates, erro = template.ParseFiles(file)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	erro = templates.Execute(w, result)
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
