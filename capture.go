package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/senghoo/web2pic/snap"
)

var (
	dockerURI string
	addr      string
)

type SnapHander struct {
}

func (h *SnapHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	s := snap.NewSnap(url, dockerURI)
	err := s.Snap()
	if err != nil {
		log.Printf("Snap %s error: %s", url, err)
	}

	defer s.Clear()

	reader, size, err := s.SnapReader()
	if err != nil {
		log.Printf("Snap %s error: %s", url, err)
	}
	defer reader.Close()

	w.Header().Set("Content-Length", fmt.Sprint(size))
	w.Header().Set("Content-Type", "image/png")

	_, err = io.Copy(w, reader)
	if err != nil {
		log.Printf("Snap %s error: %s", url, err)
	}
}

func parse() {
	flag.StringVar(&dockerURI, "docker", "unix:///var/run/docker.sock", "docker uri")
	flag.StringVar(&addr, "address", ":8080", "listen address")
}

func main() {
	// parge arguments
	parse()

	// mux
	mux := http.NewServeMux()
	mux.Handle("/snap", new(SnapHander))

	// serve
	serve := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Printf("Listening %s", serve.Addr)
	log.Fatal(serve.ListenAndServe())
}
