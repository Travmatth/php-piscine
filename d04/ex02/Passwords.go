package Ex02

import (
	Ex00 "d04/ex00"
	Ex01 "d04/ex01"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func setUser(login, password string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	Ex00.Store[login] = &Ex00.User{Login: login, Password: string(hash)}
	return Ex01.SaveUsers(Ex00.Store)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func updatePw(login, old, new string) (err error) {
	err = nil
	if current, isCurrent := Ex00.Store[login]; !isCurrent {
		err = fmt.Errorf("No password set for %s", login)
		fmt.Println(err)
	} else if !ComparePasswords(current.Password, []byte(old)) {
		err = fmt.Errorf("Passwords do no match for %s", login)
		fmt.Println(err)
	} else if err = setUser(login, new); err != nil {
		fmt.Printf("New Passwords for %s could not be saved", login)
		fmt.Println(err)
	}
	return
}

func CompareUser(w http.ResponseWriter, r *http.Request) (status string) {
	status = "ERROR\n"
	r.ParseForm()
	if submit, isSubmit := r.Form["submit"]; isSubmit && submit[0] == "OK" {
		if login, ok := r.Form["login"]; ok == false {
			fmt.Println("Login field is empty")
		} else if _, stored := Ex00.Store[login[0]]; !stored {
			fmt.Println("User does not exist")
		} else if old, exists := r.Form["oldpw"]; !exists || old[0] == "" {
			fmt.Println("Old password field is empty")
		} else if new, exists := r.Form["newpw"]; !exists || new[0] == "" {
			fmt.Println("New password field is empty")
		} else if err := updatePw(login[0], old[0], new[0]); err == nil {
			fmt.Println("User", login[0], "saved")
			status = "OK\n"
		}
	}
	return
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
