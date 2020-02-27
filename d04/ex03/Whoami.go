package Ex03

import (
	Ex00 "d04/ex00"
	"net/http"
)

/*
whoami.php wont take any arguments but will use the
session cooke to display the login contained in the
'logged_on_user' session variable followed by '\n'.
If this variable does not exist or contains an empty
string, display only "ERROR\n"
*/

func HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
	var body string
	if c, err := r.Cookie(Ex00.COOKIE_NAME); err != nil {
		body = "ERROR\n"
	} else if login, ok := Session[c.Value]; !ok {
		body = "ERROR\n"
	} else {
		body = login + "\n"
	}
	w.Write([]byte(body))
}
