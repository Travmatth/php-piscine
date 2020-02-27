package Ex02

import (
	"fmt"
	"html/template"
	"net/http"
)

var HandleIndexEndpoint string = "/j04/ex02/index.php"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	tmpl := template.Must(template.ParseFiles("ex02/index.html"))
	tmpl.Execute(w, struct{}{})
}

var HandleCreateEndpoint string = "/j04/ex02/create.php"

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	fmt.Fprintf(w, "%s", RetrieveUser(w, r))
}

var HandleModifEndpoint string = "/j04/ex02/modif.php"

func HandleModif(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	if r.Method == "POST" {
		fmt.Fprintf(w, "%s", CompareUser(w, r))
		return
	}
	tmpl := template.Must(template.ParseFiles("ex02/modify.html"))
	tmpl.Execute(w, struct{}{})
}
