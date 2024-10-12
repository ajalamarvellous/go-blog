package main

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
	"html/template"
	"first-go-project/internal/database"
)

// Home is defined as a method of application struct now
func (app *application) home(
	w http.ResponseWriter, r *http.Request){
	// Checks if the url ends with "/"
	if r.URL.Path != "/" {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}
	// Getting the most recent values in our database
	contents, err := app.db.Recent()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Contents: contents,
	}
	// We will use Execute() method to write the template content as response body
	// the second argument is supposed to be any dynamic content we want to send but 
	// leave as nil for now
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

// snippetView defined as a method of application
func (app *application) snippetView(
	w http.ResponseWriter, r *http.Request){
		// get specific id from the url query
		id, err := strconv.Atoi(r.URL.Query().Get("id")) 
		if err != nil || id < 1 {
			app.notFound(w)
			return 
		}
		content, err := app.db.Get(id)
		if err != nil {
			if errors.Is(err, database.ErrNoRecord){
				app.notFound(w)
			} else {
				app.serverError(w, err)
			}
			return
		}
		// initialising a new slice that contain list of html files including
		// new view.html
		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			"./ui/html/pages/view.html",
		}
		// parsing the files
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
		}


		err = ts.ExecuteTemplate(w, "base", content)
		if err != nil {
			app.serverError(w, err)
		}
	}

// snippetCreate as a method of application
func (app *application)snippetCreate(
	w http.ResponseWriter,
	r *http.Request){
		// if r.Method != "POST"{
		if r.Method != http.MethodPost{
			// if the method is not POST, return a 405 status and 
			// write method is not allowed
			w.Header().Set("Allow", "POST")
			app.clientError(w, http.StatusMethodNotAllowed)
			return
		}
		// creating some dummy variable for the build
		title := "Personal info"
		content := "Date of Birth: July 18, \n Hometown: Ogbomoso, Oyo state \n current profession: Jobless"
		var expires int = 90
		// pass the data into the db.insert method
		id, err := app.db.Insert(title, content, expires)
		if err != nil {
			app.serverError(w, err)
		}
		http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id),
			http.StatusSeeOther)
	}

