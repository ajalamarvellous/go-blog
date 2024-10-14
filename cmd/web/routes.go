package main

import "net/http"

func (app *application) routes() http.Handler{

	// Using http.NewServeMux to create the router
	mux := http.NewServeMux()
	// Adding a file handler to serve the frontend css and js files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/user/signup", app.userSignup)
	mux.HandleFunc("/user/login", app.userLogin)
	mux.HandleFunc("/user/logout", app.userLogoutPost)

	return secureHeaders(mux)
}