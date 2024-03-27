package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	//handelling the catch all thing
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello From SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	
	//w.Write([]byte("Hello From snippetView"))

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w,"View the specific snippet with id %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)

		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))

		//alternative
		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello From SnippetCreate"))
}