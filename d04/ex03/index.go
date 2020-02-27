package Ex03

import (
	"net/http"
	"time"

	Ex00 "d04/ex00"

	UUID "github.com/google/uuid"
)

var Session map[string]string = make(map[string]string)

func StartSession(w http.ResponseWriter, r *http.Request) (uuid string) {
	c, err := r.Cookie(Ex00.COOKIE_NAME)
	if err != nil {
		c = &http.Cookie{
			Name:    Ex00.COOKIE_NAME,
			Value:   UUID.New().String(),
			Expires: time.Unix(0, 0),
			MaxAge:  60 * 60 * 24,
			Path:    "/",
			Domain:  r.URL.Hostname(),
		}
		http.SetCookie(w, c)
	}
	uuid = c.Value
	return
}

var HandleLoginEndpoint string = "/j04/ex03/login.php"
var HandleLogoutEndpoint string = "/j04/ex03/logout.php"
var HandleWhoAmIEndpoint string = "/j04/ex03/whoami.php"
