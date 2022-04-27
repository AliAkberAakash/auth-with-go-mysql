package main

import (
	"log"
	"net/http"

	"github.com/AliAkberAakash/auth-with-go-mysql/handler"
)

func main() {

	port := ":8000"

	http.Handle("/register", handler.GetRegisterHandler())
	http.Handle("/login", handler.GetLoginHandler())

	log.Fatal(http.ListenAndServe(port, nil))
}
