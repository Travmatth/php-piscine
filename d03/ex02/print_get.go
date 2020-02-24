package main

import (
	"fmt"
	"net/http"
)

func phpInfoHandler(w http.ResponseWriter, r *http.Request) {
	str := ""
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	for key, arr := range r.URL.Query() {
		for _, val := range arr {
			str += key + ": " + val + "\n"
		}
	}
	fmt.Fprintf(w, "%s", str)
}

func main() {
	http.HandleFunc("/j03/ex02/print_get.php", phpInfoHandler)
	http.ListenAndServe(":8080", nil)
}
