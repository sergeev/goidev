package common

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"../helpers"
	"../storage"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Handlers
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Server: [net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}

// for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("tpl/login.gohtml")
	fmt.Fprintf(response, body)
}

// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		// Database check for user data!
		_userIsValid := storage.UserIsValid(name, pass)

		// Пользователь вошел в систему, перенаправляем на index
		if _userIsValid {
			SetCookie(name, response)
			redirectTarget = "/index"
		} else {
			// Не валидный вход, кидаем на регистрацию
			redirectTarget = "/register"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("tpl/register.gohtml")
	fmt.Fprintf(response, body)
}

// for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !helpers.IsEmpty(uName)
	_email = !helpers.IsEmpty(email)
	_pwd = !helpers.IsEmpty(pwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _uName && _email && _pwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

// for GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	// Проверка пользователя в базе
	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		// пользователь найден, загружаем index.gohtml
		var indexBody, _ = helpers.LoadFile("tpl/index.gohtml")
		fmt.Fprintf(response, indexBody, userName)
	} else {
		// Данного пользователя нет!
		http.Redirect(response, request, "/", 302)
	}
}

func ManageHandler(response http.ResponseWriter, request *http.Request) {
	// Проверка пользователя в базе
	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		var indexBody, _ = helpers.LoadFile("tpl/manage.gohtml")
		fmt.Fprintf(response, indexBody, userName)
	} else {
		// Данного пользователя нет!
		http.Redirect(response, request, "/", 302)
	}
}

// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	// Очищаем кэш и выходим из учетной записи
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}

// Cookie
func SetCookie(userName string, response http.ResponseWriter) {
	// грузим мап и кэш
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	// очищаем кэш
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
