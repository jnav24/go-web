package main

import (
	"github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jnav24/go-web/pkg/config"
	"github.com/jnav24/go-web/pkg/handlers"
	"net/http"
)

func useChiRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

func usePatRouter() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

func routes(app *config.AppConfig) http.Handler {
	//return usePatRouter()
	return useChiRouter()
}
