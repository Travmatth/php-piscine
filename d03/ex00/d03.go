package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	static := fmt.Sprintf("%s/http/MyWebSite/d03", os.Getenv("HOME"))
	fs := http.FileServer(http.Dir(static))
	http.Handle("/", fs)
	http.ListenAndServe(":8100", nil)
}
