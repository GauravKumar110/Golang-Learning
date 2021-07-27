package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))

	s := http.Server{
		Addr:    ":5001",
		Handler: mux,
	}

	s.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good Morning India"))
}
