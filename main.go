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



func main(){
	// Using http.NewServeMux to create the router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// starting a new server using http.ListenAndServe and port 400
	// http.ListenAndServer requires majorly 2 parameters, port and router
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// if any error (http.ListenAndServe returns non nil), log error and close
	log.Fatal(err)
}