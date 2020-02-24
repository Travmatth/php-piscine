package main

import (
	"fmt"
	"net/http"
)

func rawTextHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "<html><body>Hello</body></html>")
}

func main() {
	http.HandleFunc("/ex04/raw_text.php", rawTextHandler)
	http.ListenAndServe(":8080", nil)
}
