package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the firstpage")
	fmt.Println("Endpoint")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
