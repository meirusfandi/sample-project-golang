package main

import (
	"fmt"
	"net/http"
	"../sampleproject/handler"
)

func main() {
	http.HandleFunc("/index", handler.IndexUserPage)
	http.HandleFunc("/index/user", handler.IndexOneUserPage)
	http.HandleFunc("/index/add", handler.AddUserPage)
	http.HandleFunc("/index/update", handler.UpdateUserPage)
	http.HandleFunc("/index/delete", handler.DeleteUserPage)

	fmt.Println("server started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}