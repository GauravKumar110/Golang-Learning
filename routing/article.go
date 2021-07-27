package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the firstpage")
	fmt.Println("Endpoint")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)

}

func main() {
	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()
}
