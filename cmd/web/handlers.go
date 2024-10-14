package main

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
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

	// Getting the most recent values in our database
	contents, err := app.db.Recent()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Contents: contents,
	}
	app.render(w, http.StatusOK, "home.html", data)

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
		
		data := &templateData{
			Content: content,
		}

		app.render(w, http.StatusOK, "view.html", data)
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


func (app *application) userSignup (
	w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintln(w, "Display the html form for signing up new users...")
		} else {
			if r.Method == "POST" {
				app.userSignup(w, r)
			}
		}
	
	}

func (app *application) userSignupPost(
	w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create a new user...")
	}

func (app *application) userLogin(
	w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost{
			app.userLoginPost(w,r)
		} else {
			if r.Method == http.MethodGet {
				fmt.Fprintln(w, "Display the login page here...")
			}
		}
	}

func (app *application) userLoginPost(
	w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Authenticating user...")
	}

func (app *application) userLogoutPost(
	w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Logging out user...")
	}