package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist/", artist)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Web server is on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
