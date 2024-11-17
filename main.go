package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Print("Starting server on localhost:8080")

	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
