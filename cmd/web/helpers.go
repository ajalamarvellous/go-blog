package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(
	w http.ResponseWriter, err error){
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.errorLog.Println(trace)
		http.Error(
			w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

func (app *application) clientError(
	w http.ResponseWriter, status int){
		http.Error(w, http.StatusText(status), status)
	}

func (app *application) notFound(w http.ResponseWriter){
		app.clientError(w, http.StatusNotFound)
	}

func (app *application) render(
	w http.ResponseWriter, status int, page string, data *templateData){
		// check if the template we want to call exist in our map
		ts, ok := app.templateCache[page]
		if !ok {
			err := fmt.Errorf("The template does not exist", page)
			app.serverError(w, err)
		}

		// initialising a new buffer content
		buf := new(bytes.Buffer)

		
		// execute the template set and write response in the body
		// err := ts.ExecuteTemplate(w, "base", data)
		err := ts.ExecuteTemplate(buf, "base", data)
		if err != nil {
			app.serverError(w, err)
		}

		// write out the provided HTTP status code (200-ok, 400-Bad request)
		w.WriteHeader(status)

		// we write the content of buff to the http.ResponseWriter
		buf.WriteTo(w)
	}