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

	//handler function 

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {

	//handelling the catch all thing
	if(r.URL.Path != "/"){
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello From SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From snippetView"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From SnippetCreate"))
}
