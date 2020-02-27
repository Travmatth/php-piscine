package Ex01

import (
	Ex00 "d04/ex00"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func PreparePasswordFile() (err error) {
	dir := fmt.Sprintf("%s/http/MyWebSite/j04/htdocs/private", os.Getenv("HOME"))
	PATH = dir
	FILE = dir + "/passwd"
	if _, err = os.Stat(PATH); os.IsNotExist(err) {
		if os.Mkdir(PATH, os.ModePerm) != nil {
			fmt.Println("Error opening directory:", err)
			return
		}
	}
	if _, err = os.Stat(FILE); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(FILE); err != nil {
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
	Ex00.Store[login] = &Ex00.User{Login: login, Password: string(hash)}
	return SaveUsers(Ex00.Store)
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) (status string) {
	status = "ERROR\n"
	r.ParseForm()
	if submit, isSubmit := r.Form["submit"]; isSubmit && submit[0] == "OK" {
		if login, ok := r.Form["login"]; ok == false {
			fmt.Println("Login field is empty")
		} else if _, stored := Ex00.Store[login[0]]; stored {
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
