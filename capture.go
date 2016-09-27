package main

import (
	"log"
	"net/http"
)

type SnapHander struct {
}

func (h *SnapHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/snap", new(SnapHander))
	serve := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(serve.ListenAndServe())
}
