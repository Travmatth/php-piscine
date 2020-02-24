package main

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var html = "<html><body>Hello Zaz<br/><img src='data:image/png;base64,IMAGE'></body></html>"

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	image, err := ioutil.ReadFile("42.png")
	if err != nil {
		log.Fatal(err)
	}
	encoded := base64.StdEncoding.EncodeToString(image)
	body := strings.Replace(html, "IMAGE", encoded, 1)
	w.Header().Set("Content-Type", "image/png")
	w.Write([]byte(body))
}

var username string = "zaz"
var password string = "jaimelespetitsponeys"

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		ok = !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) == 1
		ok = !ok || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"Member area\"")
			w.WriteHeader(401)
			w.Write([]byte("<html><body>That area is accessible for members only</body></html>"))
			return
		}
	}
	return handler(w, r)
}

func main() {
	http.HandleFunc("/ex04/raw_text.php", BasicAuth(ResponseHandler))
	http.ListenAndServe(":8080", nil)
}
