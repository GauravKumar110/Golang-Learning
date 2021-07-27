package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int
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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	intVar, err := strconv.Atoi(key)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	fmt.Fprintf(w, "Key: "+key)

	for _, article := range Articles {
		if article.Id == intVar {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()
}
