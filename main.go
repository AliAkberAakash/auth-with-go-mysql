package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AliAkberAakash/auth-with-go-mysql/controller"
	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"github.com/AliAkberAakash/auth-with-go-mysql/response"
)

func main() {

	port := ":8000"

	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/register", registerHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var userCredential model.User

	err := json.NewDecoder(r.Body).Decode(&userCredential)

	if err != nil {
		json.NewEncoder(w).Encode(response.GetErrorResponse(http.StatusBadRequest, "Invalid Request Body"))
		return
	}

	controller.Register(userCredential)

	json.NewEncoder(w).Encode(controller.Users)
	return
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
