package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"../database"
)

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
	var result = database.GetAllUser()
	templates, err = templates.ParseGlob("view/*.html")
	checkErr(err)

	for _, each := range result {
		fmt.Println(each)
	}

	//replacing data to template index
	err = templates.ExecuteTemplate(w, "index.html", result)
	if err != nil {
		fmt.Fprintln(w, "Error renderin page"+err.Error())
	}
}

func IndexOneUserPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	var result = database.GetUser(idx)
	err = templates.ExecuteTemplate(w, "user.html", result)
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

	id := r.URL.Query().Get("id")

	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	var result = database.GetUser(idx)

	err = templates.ExecuteTemplate(w, "update.html", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to delete page")
}

func ActionAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var uname = r.FormValue("username")
		var pass = r.FormValue("password")
		var email = r.FormValue("email")
		var name = r.FormValue("fullname")

		var user = database.User{}

		user.Username = uname
		user.Fullname = name
		user.Email = email
		user.Password = pass

		database.AddUser(user)
	}

	http.Redirect(w, r, "/index", 301)
}

func ActionUpdateUser(w http.ResponseWriter, r *http.Request) {

}
