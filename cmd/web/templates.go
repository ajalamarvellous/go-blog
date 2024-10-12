package main

import "first-go-project/internal/database"

// defining a new templatedata type for holding dynamic data
// we want to pass into our html pages

type templateData struct {
	Content *database.Content
	Contents []*database.Content
}