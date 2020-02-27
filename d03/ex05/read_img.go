package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var buf *io.Reader

func imageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	w.Header().Set("Content-Type", "image/png")
	image, err := ioutil.ReadFile("../img/42.png")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(image)
}

func main() {
	http.HandleFunc("/ex05/read_img.php", imageHandler)
	http.ListenAndServe(":8080", nil)
}
