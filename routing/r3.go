package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)

	s := http.Server{
		Addr:    ":5002",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Haridwar"))
}
