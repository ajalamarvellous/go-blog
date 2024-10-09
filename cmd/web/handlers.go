package main

import (
	"fmt"
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

