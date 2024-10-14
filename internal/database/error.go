package database

import "errors"

var ErrNoRecord = errors.New("No matching record found")

// Adding an invalid creditials error
var ErrInvalidCredentials = errors.New("models: invalid credentials")

// Adding a duplicated email error
var ErrDuplicatedEmail = errors.New("models: duplicated email")