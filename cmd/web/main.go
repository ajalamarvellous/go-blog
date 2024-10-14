package main

import (
	"os"
	"flag"
	"log"
	"net/http"
	"database/sql"
	"html/template"
	"first-go-project/internal/database"  // the db model we defined 
    _ "github.com/go-sql-driver/mysql"
)


// creating an application struct with dependencies 
// will only contain the different loggers for a start
type application struct{
	errorLog 	*log.Logger
	infoLog 	*log.Logger
	users		*database.UserDB
	db 			*database.DatabaseModel
	templateCache  map[string]*template.Template
}

func main(){
	// define commandline arguments for port-address
	// the arguments in order name_of_arg, default_value, description
	addr := flag.String("addr", ":4000", "HTTP Port to use")
	// Add a flag for the database manager
    dsn := flag.String("dsn", "web:pass@/go_project?parseTime=true", "MySQL data source name")
	flag.Parse()


	// creating two different loggers for info and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// instantiating the database pool 
    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    // Ensure every other function has ran before runinng db.Close
    // but also db closes before the main file closes
    defer db.Close()

	// initiating the new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// instantiating application class
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		db: &database.DatabaseModel{DB: db},
		users : &database.UserDB{DB: db},
		templateCache: templateCache,
	}


	// instantiating a new server object that takes in our port address
	// error logger and server handler
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	// starting a new server using http.ListenAndServe and port 400
	// http.ListenAndServer requires majorly 2 parameters, port and router
	app.infoLog.Printf("Starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

	// if any error (http.ListenAndServe returns non nil), log error and close
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    return db, nil
}