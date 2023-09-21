package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave)

	// register routes
	mux.Get("/", app.home)
	mux.Post("/login", app.login)

	mux.Route("/user", func(r chi.Router) {
		r.Use(app.auth)
		r.Get("/profile", app.profile)
	})

	// static assets
	fileserver := http.FileServer(http.Dir("./webapp/static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
