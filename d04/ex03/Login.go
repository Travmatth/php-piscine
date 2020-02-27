package Ex03

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

/*
Create login.php that takes login and password  variables in the url.
Should start the session, check the validity of the combo login/password
combo and store in the session a logged_on_user that contains login of
the user that has submitted the correct combo, or an empty string otherwise
*/

func GetLoginAndPassword(v url.Values) (login, password string, err error) {
	if login == "" || password == "" {
		err = errors.New("Login and Password not found")
	}
	return
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	sessionUUID := StartSession(w, r)
	login, password := v.Get("login"), v.Get("passwd")
	fmt.Println("Login reached", login, password, sessionUUID)
	if login != "" && password != "" && Auth(login, password, sessionUUID) {
		fmt.Println("Login successful")
		Session[sessionUUID] = login
		w.Write([]byte("OK\n"))
		return
	}
	fmt.Println("Login failed")
	w.Write([]byte("ERROR\n"))
}
