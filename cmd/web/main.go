package main

import (
	"os"
	"flag"
	"log"
	"net/http"
)


// creating an application struct with dependencies 
// will only contain the different loggers for a start
type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}

func main(){
	// define commandline arguments for port-address
	// the arguments in order name_of_arg, default_value, description
	addr := flag.String("addr", ":4000", "HTTP Port to use")
	flag.Parse()

	
	// creating two different loggers for info and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// instantiating application class
	app := application{
		errorLog: errorLog,
		infoLog: infoLog,
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
	err := srv.ListenAndServe()

	// if any error (http.ListenAndServe returns non nil), log error and close
	errorLog.Fatal(err)
}