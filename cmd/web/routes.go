package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iamlego/bookingBoost/pkg/config"
	handler "github.com/iamlego/bookingBoost/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/About", http.HandlerFunc(handler.Repo.About))
	mux := chi.NewRouter()

	mux.Use(WriteToDConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/About", handler.Repo.About)
	mux.Get("/generals-quarters", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Get("/make-reservation", handler.Repo.Reservation)
	mux.Get("/contact", handler.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	// log.Println("The content of file server are : ", fileServer) // it is &{./static}

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
