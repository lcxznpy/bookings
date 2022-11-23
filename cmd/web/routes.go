package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lcxznpy/bookings/pkg/config"
	"github.com/lcxznpy/bookings/pkg/handles"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handles.Repo.Home)
	mux.Get("/about", handles.Repo.About)
	mux.Get("/one", handles.Repo.One)
	mux.Get("/two", handles.Repo.Two)
	mux.Get("/contact", handles.Repo.Contact)
	mux.Get("/search-availability", handles.Repo.Availability)
	mux.Get("/make-reservation", handles.Repo.Reservation)

	//静态图像的显示
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
