package main

import (
	"html/template"
	"path/filepath"

	"first-go-project/internal/database"
)
// defining a new templatedata type for holding dynamic data
// we want to pass into our html pages

type templateData struct {
	Content *database.Content
	Contents []*database.Content
}

func newTemplateCache() (map[string]*template.Template, error){
	// initializing a new map (dictionary) to serve as our cache
	cache := map[string]*template.Template{}

	// getting a list of all the html pages in html/pages
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	// looping throught the pages
	for _, page := range pages {
		// Extract eh file name from the file path
		name := filepath.Base(page)
	
		files := []string {
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			page,
		}
		// parse the files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// caching the template set to map
		cache[name] = ts
	}
	return cache, nil
}

