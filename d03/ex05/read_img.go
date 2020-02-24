package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	image, err := os.Open("42.png")
	if err != nil {
		log.Fatal(err)
	}
	defer image.Close()
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, image)
}

func main() {
	http.HandleFunc("/ex04/raw_text.php", imageHandler)
	http.ListenAndServe(":8080", nil)
}
