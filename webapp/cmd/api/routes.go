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
	mux.Use(app.enableCORS)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./webapp/html/"))))

	mux.Route("/web", func(r chi.Router) {
		r.Post("/auth", app.authenticate)
		r.Get("/refresh-token", app.refreshUsingCookie)
		r.Get("/logout", app.deleteRefreshCookie)
	})

	// authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// protected routes
	mux.Route("/users", func(r chi.Router) {
		r.Use(app.authRequired)

		r.Get("/", app.allUsers)
		r.Get("/{userID}", app.getUser)
		r.Delete("/{userID}", app.deleteUser)
		r.Put("/", app.insertUser)
		r.Patch("/", app.updateUser)
	})

	return mux
}
