package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":8000"

	http.HandleFunc("/", helloWorldHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
