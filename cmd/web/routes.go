package main

import (
	"net/http"

	"github.com/cpratap994/bookings-learngo/pkg/config"
	"github.com/cpratap994/bookings-learngo/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/luxury-room", handlers.Repo.LuxuryRoom)
	mux.Get("/president-suite", handlers.Repo.PresidentSuite)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
