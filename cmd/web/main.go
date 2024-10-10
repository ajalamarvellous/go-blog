package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){
	// define commandline arguments for port-address
	// the arguments in order name_of_arg, default_value, description
	addr := flag.String("addr", ":4000", "HTTP Port to use")
	flag.Parse()

	// Using http.NewServeMux to create the router
	mux := http.NewServeMux()
	// Adding a file handler to serve the frontend css and js files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// starting a new server using http.ListenAndServe and port 400
	// http.ListenAndServer requires majorly 2 parameters, port and router
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	// if any error (http.ListenAndServe returns non nil), log error and close
	log.Fatal(err)
}