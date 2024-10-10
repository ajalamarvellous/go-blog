package main

import (
	"log"
	"fmt"
	"strconv"
	"net/http"
	"html/template"
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
	// We will link the html template page with the home handler
	// using template.ParseFiles(), if there's error, we log and send generic 500 to user
	// We will also initialise a slice containing both our base and specific page with the 
	// base file being the first file to mention/list
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// We will use Execute() method to write the template content as response body
	// the second argument is supposed to be any dynamic content we want to send but 
	// leave as nil for now
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Error occurred", 500)
	}
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

