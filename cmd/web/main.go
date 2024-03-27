package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Hello World")

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	//handler function

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}