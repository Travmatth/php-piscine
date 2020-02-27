package Ex01

import (
	"fmt"
	"html/template"
	"net/http"
)

var PATH string
var FILE string

var HandleHomeEndpoint string = "/j04/ex01/index.php"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	tmpl := template.Must(template.ParseFiles("ex01/index.html"))
	tmpl.Execute(w, struct{}{})
}

var HandleCreateEndpoint string = "/j04/ex01/create.php"

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	if r.Method == "POST" {
		fmt.Fprintf(w, "%s", RetrieveUser(w, r))
	}
}
