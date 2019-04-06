package main

import (
	common "./common"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router = mux.NewRouter()

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		log.Printf("Server: [net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}


func main() {
	// Загрузка стилей проекта
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	router.HandleFunc("/", Logger(common.LoginPageHandler)) // GET

	router.HandleFunc("/index", Logger(common.IndexPageHandler)) // GET
	router.HandleFunc("/login", Logger(common.LoginHandler)).Methods("POST")

	router.HandleFunc("/register", Logger(common.RegisterPageHandler)).Methods("GET")
	router.HandleFunc("/register", Logger(common.RegisterHandler)).Methods("POST")

	router.HandleFunc("/logout", Logger(common.LogoutHandler)).Methods("POST")

	// Manage page (admin)
	router.HandleFunc("/manage", Logger(common.ManageHandler)) // GET

	http.Handle("/", router)
	log.Print("Server: start")
	http.ListenAndServe(":8080", nil)
}