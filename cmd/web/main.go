package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	

	//handler function
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
