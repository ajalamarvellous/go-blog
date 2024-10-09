package main

import (
	"fmt"
	"log"
	"strconv"
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
		// get specific id from the url query
		id, err := strconv.Atoi(r.URL.Query().Get("id")) 
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return 
		}
		fmt.Fprintf(w, "Displat this specific response with ID %d..", id)
	}


func snippetCreate(
	w http.ResponseWriter,
	r *http.Request){
		// if r.Method != "POST"{
		if r.Method != http.MethodPost{
			// if the method is not POST, return a 405 status and 
			// write method is not allowed
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method is not Allowed", http.StatusMethodNotAllowed)
			return
		}
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