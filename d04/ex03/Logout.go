package Ex03

import (
	Ex00 "d04/ex00"
	"net/http"
)

/*
logout.php wont take any arguments, but uses the
session cookie to remove the last one. Page will display
nothing.
*/

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(Ex00.COOKIE_NAME)
	if err != nil {
		return
	}
	_, ok := Session[c.Value]
	if !ok {
		return
	}
	delete(Session, c.Value)
}
