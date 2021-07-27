package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := new(TestHandler)
	mux.Handle("/", handler)

	s := http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	s.ListenAndServe()
}

type TestHandler struct{}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome India"))
}
