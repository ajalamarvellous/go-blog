package database

import (
	"errors"
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
	// SQL query to get specific data id
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE id=?;"
	// using exec to execute the query
	row := db.DB.QueryRow(stmt, id)
	
	// creating a new pointer for the content row 
	s := &Content{}

	// using row scan to copy the content of the row to the content object
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// this will return the 10 most recent contents
func (db *DatabaseModel) Recent() ([]*Content, error){
	// SQL query to get specific data id
	stmt := "SELECT id, title, content, created, expires FROM snippets LIMIT 10;"
	// using exec to execute the query
	rows, err := db.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// defering the closing of the row
	defer rows.	Close()
	
	// creating an array to contain all our data
	data := []*Content{}

	//interating through the rows
	for rows.Next(){
		// creating a new pointer for the content row 
		s := &Content{}

		// using row scan to copy the content of the row to the content object
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		
		if err != nil {
			return nil, err
		}
		data = append(data, s)
	}
		if err = rows.Err(); err != nil {
			return nil, err
		} 
		// errors.Is(err, sql.ErrNoRows){
		// 	return nil, ErrNoRecord
		// } else {
		// 	return nil, err
		// }
	
	return data, nil
}
