package main

import (
	Ex00 "d04/ex00"
	Ex01 "d04/ex01"
	Ex02 "d04/ex02"
	Ex03 "d04/ex03"
	"fmt"
	"net/http"
)

func main() {
	if err := Ex01.PreparePasswordFile(); err != nil {
		return
	}
	http.HandleFunc(Ex00.HandleAuthenticationEndpoint, Ex00.HandleAuthentication)
	http.HandleFunc(Ex01.HandleHomeEndpoint, Ex01.HandleHome)
	http.HandleFunc(Ex01.HandleCreateEndpoint, Ex01.HandleCreate)
	http.HandleFunc(Ex02.HandleIndexEndpoint, Ex02.HandleIndex)
	http.HandleFunc(Ex02.HandleCreateEndpoint, Ex02.HandleCreate)
	http.HandleFunc(Ex02.HandleModifEndpoint, Ex02.HandleModif)
	http.HandleFunc(Ex03.HandleLoginEndpoint, Ex03.HandleLogin)
	http.HandleFunc(Ex03.HandleLogoutEndpoint, Ex03.HandleLogout)
	http.HandleFunc(Ex03.HandleWhoAmIEndpoint, Ex03.HandleWhoAmI)
	fmt.Println("Listening at 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
