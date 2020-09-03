package main

import (
	"database/sql"
	"log"
	"net/http"

	common "./common"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

var router = mux.NewRouter()

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		log.Printf("Server: [net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}

func main() {
	// Подключение к базе данных
	db, err = sql.Open("mysql", "root:root@tcp(localhost:8889)/golang_auth")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Загрузка стилей проекта
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Главная страница проекта
	router.HandleFunc("/", Logger(common.WelcomePageHandler)) // GET

	router.HandleFunc("/contacts", Logger(common.ContactsPageHandler)) // GET

	//router.HandleFunc("/", Logger(common.LoginPageHandler)) // GET

	router.HandleFunc("/index", Logger(common.IndexPageHandler)) // GET
	router.HandleFunc("/login", Logger(common.LoginPage)).Methods("POST")

	//router.HandleFunc("/register", Logger(common.RegisterPageHandler)).Methods("GET")
	router.HandleFunc("/register", Logger(common.SignupPage)).Methods("GET")
	router.HandleFunc("/register", Logger(common.SignupPage)).Methods("POST")

	router.HandleFunc("/logout", Logger(common.LogoutHandler)).Methods("POST")

	// Manage page (admin)
	router.HandleFunc("/manage", Logger(common.ManageHandler)) // GET

	http.Handle("/", router)
	log.Print("Server: start")
	http.ListenAndServe(":8080", nil)

}
