package main

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/vfaust1/movie-api/docs"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /movies", app.getAllMoviesHandler)
	router.HandleFunc("GET /movies/{id}", app.getMovieByIDHandler)
	router.HandleFunc("POST /movies", app.createMovieHandler)
	router.HandleFunc("PUT /movies/{id}", app.updateMovieHandler)
	router.HandleFunc("DELETE /movies/{id}", app.deleteMovieHandler)

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	return app.loggingMiddleware(app.authMiddleware(router))
}
