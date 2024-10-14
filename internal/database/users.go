package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID 			int
	Name 		string
	Email		string
	Password 	[]byte
	created 	time.Time
}

type UserDB struct {
	DB *sql.DB
}

// Insert function to add to the users.DB
func (m *UserDB) Insert(name, email, password string) error {
	return nil
}

// Aunthenticate function to verify users
func (m *UserDB) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exist function to check if user exist
func (m *UserDB) Exist(id int) (bool, error) {
	return false, nil
}