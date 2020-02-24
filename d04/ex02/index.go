package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Login    string
	Password string
}

var path = "../private"

var file = "../private/passwd"

var cookieName = "SessionID"

var store = make(map[string]*User, 0)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, struct{}{})
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	fmt.Fprintf(w, "%s", RetrieveUser(w, r))
}

func handleModif(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	if r.Method == "POST" {
		fmt.Fprintf(w, "%s", CompareUser(w, r))
		return
	}
	tmpl := template.Must(template.ParseFiles("modify.html"))
	tmpl.Execute(w, struct{}{})
}

func main() {
	if PreparePasswordFile() == nil {
		http.HandleFunc("/j04/ex02/index.php", handleIndex)
		http.HandleFunc("/j04/ex02/create.php", handleCreate)
		http.HandleFunc("/j04/ex02/modif.php", handleModif)
		fmt.Println("Listening at 0.0.0.0:8080")
		http.ListenAndServe(":8080", nil)
	}
}
