package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)
	mux.Get("/movies", app.AllMovies)

	return mux
}
