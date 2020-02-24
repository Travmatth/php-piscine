package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	UUID "github.com/google/uuid"
)

var COOKIE_NAME = "SessionID"
var store = make(map[string]*User, 0)

type User struct {
	Login    string
	Password string
}

func makeCookie(r *http.Request) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   UUID.New().String(),
		Expires: time.Unix(0, 0),
		MaxAge:  60 * 60 * 24,
		Path:    "/",
		Domain:  r.URL.Hostname(),
	}
	return
}

func checkLogin(r *http.Request, user *User) {
	login, loginOk := r.Form["login"]
	if !loginOk {
		login[0] = ""
	}
	password, passwordOk := r.Form["passwd"]
	if !passwordOk {
		password[0] = ""
	}
	user.Login = login[0]
	user.Password = password[0]
}

func retrieveUser(domain string, r *http.Request) (user *User) {
	r.ParseForm()
	submit, isSubmit := r.Form["submit"]
	submitted := isSubmit && submit[0] == "OK"
	user, userExists := store[domain]

	switch {
	case !userExists && !isSubmit:
		user = &User{Login: "", Password: ""}
		store[domain] = user
	case userExists && !isSubmit:
	case !userExists && submitted:
		user = &User{Login: "", Password: ""}
		store[domain] = user
		fallthrough
	case userExists && submitted:
		checkLogin(r, user)
	}
	fmt.Printf("Setting user for domain: %s, %+v\n", domain, user)
	return
}

func handleAuthentication(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	tmpl := template.Must(template.ParseFiles("index.html"))
	cookie, err := r.Cookie(COOKIE_NAME)
	user := &User{Login: "", Password: ""}
	if err != nil {
		cookie = makeCookie(r)
		http.SetCookie(w, cookie)
		fmt.Println("Setting cookie", cookie.Value, "for domain", cookie.Domain)
	} else {
		fmt.Println("Detected cookie", cookie.Value, "for domain", cookie.Domain)
		user = retrieveUser(r.URL.Hostname(), r)
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/j04/ex00/index.php", handleAuthentication)
	fmt.Println("Listening at 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
