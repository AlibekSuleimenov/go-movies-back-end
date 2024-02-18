package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	Domain string
}

func main() {
	// set app config
	var app Application
	app.Domain = "example.com"

	// read from command line

	// connect to the database

	// start a web server
	log.Println("Starting app on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
