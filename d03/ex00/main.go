package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Responding to request: %s\n", r.RemoteAddr)
	name := r.URL.Path[1:]
	if len(name) == 0 {
		name = "No Name (to change, navigate to localhost:8080/<name>"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}
