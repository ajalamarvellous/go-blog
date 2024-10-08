package main

import (
	"log"
	"net/http"
)

// Define home handler func, write a byte slice 
// containing "Hello from Marve" as response
func home(
	w http.ResponseWriter, r *http.Request){
	// Checks if the url ends with "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Marve"))
}

func snippetView(
	w http.ResponseWriter, r *http.Request){
		w.Write([]byte("View the snippet..."))
	}


func snippetCreate(
	w http.ResponseWriter,
	r *http.Request){
		w.Write([]byte("Creating a snippet"))
	}

func main(){
	// Using http.NewServeMux to create the router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// starting a new server using http.ListenAndServe and port 400
	// http.ListenAndServer requires majorly 2 parameters, port and router
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// if any error (http.ListenAndServe returns non nil), log error and close
	log.Fatal(err)
}