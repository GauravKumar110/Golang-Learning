package method

import (
	"log"
	"net/http"

	comment "Golang-Learning/mini_project/controller/comment"
	post "Golang-Learning/mini_project/controller/post"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/blog", comment.Insert).Methods("POST")
	myRouter.HandleFunc("/blog", comment.GetResult).Methods("GET")
	myRouter.HandleFunc("/post", post.Create).Methods("POST")
	myRouter.HandleFunc("/post", post.GetRecords).Methods("GET")
	myRouter.HandleFunc("/post", post.DeleteRecord).Methods("DELETE")
	myRouter.HandleFunc("/post", post.UpdateRecord).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
