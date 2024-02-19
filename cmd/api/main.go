package main

import (
	"flag"
	"fmt"
	"github.com/alibeksuleimenov/go-movies-back-end/internal/repository"
	"github.com/alibeksuleimenov/go-movies-back-end/internal/repository/dbrepo"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	// set app config
	var app Application
	app.Domain = "example.com"

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=54325 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	// start a web server
	log.Println("Starting app on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
