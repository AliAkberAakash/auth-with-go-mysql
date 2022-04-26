package main

import (
	"log"
	"net/http"

	"github.com/AliAkberAakash/auth-with-go-mysql/handler"
)

func main() {

	port := ":8000"

	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/login", handler.LoginHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}
