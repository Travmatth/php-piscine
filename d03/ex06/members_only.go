package main

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var username string = "zaz"
var password string = "jaimelespetitsponeys"
var accepted string
var denied string = "<html><body>That area is accessible for members only</body></html>"

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(accepted))
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("X-Powered-By", "PHP/1.0.0")
		w.Header().Set("Server", "Apache")
		user, pass, ok := r.BasicAuth()
		userMatch := subtle.ConstantTimeCompare([]byte(user), []byte(username)) == 1
		passMatch := subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1
		if !ok || !userMatch || !passMatch {
			w.Header().Set("HTTP/1.0", "401 Unauthorized")
			w.Header().Set("WWW-Authenticate", "Basic realm=\"Member area\"")
			w.WriteHeader(401)
			w.Write([]byte(denied))
			fmt.Printf("Request from %s to %s failed\n", r.RemoteAddr, r.RequestURI)
			fmt.Println("user", userMatch, "pass", passMatch)
			return
		}
		handler(w, r)
	}
}

func main() {
	image, err := ioutil.ReadFile("../img/42.png")
	if err != nil {
		log.Fatal(err)
	}
	encoded := base64.StdEncoding.EncodeToString(image)
	accepted = "<html><body>Hello Zaz<br/><img src='data:image/png;base64,%s'></body></html>"
	accepted = fmt.Sprintf(accepted, encoded)
	http.HandleFunc("/ex06/members_only.php", BasicAuth(ResponseHandler))
	http.ListenAndServe(":8080", nil)
}
