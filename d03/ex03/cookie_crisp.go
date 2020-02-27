package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var cookies = make(map[string]http.Cookie, 0)

func setCookie(w http.ResponseWriter, name []string, nameOk bool, value []string, valueOk bool) {
	if nameOk && valueOk {
		expires := time.Now().Add(time.Hour * time.Duration(60))
		cookie := http.Cookie{Name: name[0], Value: value[0], Expires: expires}
		http.SetCookie(w, &cookie)
		fmt.Println("Cookie set for", cookie.Name)
	} else {
		fmt.Println("Error setting cookie: name", nameOk, "value", valueOk)
	}
}

func getCookie(w http.ResponseWriter, r *http.Request, name []string, nameOk bool) {
	if nameOk {
		cookie, err := r.Cookie(name[0])
		if err == nil {
			fmt.Fprintf(w, "%s\n", cookie.Value)
			fmt.Println("Cookie retrieved:", cookie)
		} else {
			fmt.Println("Error getting Cookie:", err)
		}
	}
}

func delCookie(w http.ResponseWriter, name []string, nameOk bool) {
	if nameOk {
		cookie := http.Cookie{Name: name[0], Value: "", Expires: time.Unix(0, 0)}
		http.SetCookie(w, &cookie)
		fmt.Println("Deleted Cookie: ", cookie.Name)
	} else {
		fmt.Println("Cookie cannot be deleted: Not found")
	}
}

func cookieCrispHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	params := r.URL.Query()
	if action, ok := params["action"]; ok {
		name, nameOk := params["name"]
		value, valueOk := params["value"]
		switch strings.ToLower(action[0]) {
		case "set":
			setCookie(w, name, nameOk, value, valueOk)
		case "get":
			getCookie(w, r, name, nameOk)
		case "del":
			delCookie(w, name, nameOk)
		default:
			fmt.Println("Unrecognized action requested:", action[0])
		}
	}
}

func main() {
	http.HandleFunc("/ex03/cookie_crisp.php", cookieCrispHandler)
	http.ListenAndServe(":8080", nil)
}
