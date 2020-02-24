package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login    string
	Password string
}

var path = "../private"
var file = "../private/passwd"
var cookieName = "SessionID"
var store = make(map[string]*User, 0)

func preparePasswordFile() (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if os.Mkdir(path, os.ModePerm) != nil {
			fmt.Println("Error opening directory:", err)
			return
		}
	}
	if _, err = os.Stat(file); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(file); err != nil {
			fmt.Println("Error creating password file:", err)
		}
		if err = f.Close(); err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}
	return
}

func setUser(login, password string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	store[login] = &User{Login: login, Password: string(hash)}
	return SaveUsers(store)
}

func retrieveUser(w http.ResponseWriter, r *http.Request) (status string) {
	status = "ERROR\n"
	r.ParseForm()
	if submit, isSubmit := r.Form["submit"]; isSubmit && submit[0] == "OK" {
		if login, ok := r.Form["login"]; ok == false {
			fmt.Println("Login field is empty")
		} else if _, stored := store[login[0]]; stored {
			fmt.Println("User already exists")
		} else if pw, exists := r.Form["passwd"]; !exists || pw[0] == "" {
			fmt.Println("Password field is empty")
		} else if err := setUser(login[0], pw[0]); err == nil {
			fmt.Println("User", login[0], "saved")
			status = "OK\n"
		}
	}
	return
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, struct{}{})
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	if r.Method == "POST" {
		fmt.Fprintf(w, "%s", retrieveUser(w, r))
	}
}

func main() {
	if preparePasswordFile() == nil {
		http.HandleFunc("/j04/ex01/index.php", handleHome)
		http.HandleFunc("/j04/ex01/create.php", handleCreate)
		fmt.Println("Listening at 0.0.0.0:8080")
		http.ListenAndServe(":8080", nil)
	}
}
