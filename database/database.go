package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type AllUser []User

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}