package main

import (
	"net/http"
	"udemy/pkg/config"
	"udemy/pkg/handlers"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(WriteConsole)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
