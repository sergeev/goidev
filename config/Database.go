package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB
var err error

func DatabaseConfig(response http.Response, request *http.Request) {
	// Load database
	// db, err = sql.Open("mysql", "myUsername:myPassword@/myDatabase")
	db, err = sql.Open("mysql", "root:10184902125410@/golang_db")
	if err != nil {
	panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
	panic(err.Error())
	}
}