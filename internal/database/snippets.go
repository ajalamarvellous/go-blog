package database

import (
	"database/sql"
	"time"
)

// define a content class/struct
type Content struct{
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

// defining a database model
type DatabaseModel struct{
	DB *sql.DB
}

// function to insert doc into the database
func (db *DatabaseModel) Insert(
	title string, content string, expires int) (int, error){
		// we want to enter values into our database
		stmt := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL? DAY));"
		// using Exec to execute the prompt
		result, err := db.DB.Exec(stmt, title, content, expires)
		if err != nil {
			return 0, err
		}
		// using a mysql custom method LastInsertId() to get the most recent insert
		id, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(id), nil
	}

// function to get a particular id
func (db *DatabaseModel) Get(id int) (*Content, error){
	return nil, nil
}

// this will return the 10 most recent contents
func (db *DatabaseModel) Recent() ([]*Content, error){
	return nil, nil
}