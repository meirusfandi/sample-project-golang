package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"../database"
)

// var templates = template.Must(template.ParseGlob("view/*"))
var templates *template.Template
var err error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//index for get all user on one page
func IndexUserPage(w http.ResponseWriter, r *http.Request) {

	//get data from database
	// var result []database.User
	var result = database.GetAllUser()
	templates, err = templates.ParseGlob("view/*.html")
	checkErr(err)

	for _, each := range result {
		fmt.Println(each)
	}

	//replacing data to template index
	templates.ExecuteTemplate(w, "index.html", result)
}

func IndexOneUserPage(w http.ResponseWriter, r *http.Request) {

	err = templates.ExecuteTemplate(w, "user.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUserPage(w http.ResponseWriter, r *http.Request) {

	err = templates.ExecuteTemplate(w, "add.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUserPage(w http.ResponseWriter, r *http.Request) {

	err = templates.ExecuteTemplate(w, "update.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to delete page")
}
