package main

import (
	"fmt"
	"net/http"
	"runtime"
)

var file = `<!DOCTYPE html>
<html>
	<head>
		<title>phpinfo()</title>
		<style>
			body {
				text-align: center;
			}
		</style>
    </head>
    <body bgcolor="lightgrey">
		<h1>phpinfo() - or the golang equivalent</h1>
		<hr/>
        <p>%s</p>
    </body>
</html>`

func phpInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)
	fmt.Fprintf(w, file, runtime.Version())
}

func main() {
	http.HandleFunc("/ex01/phpinfo.php", phpInfoHandler)
	http.ListenAndServe(":8080", nil)
}
