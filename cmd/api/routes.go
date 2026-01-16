package main

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/vfaust1/movie-api/docs"
)

func routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /movies", getAllMoviesHandler)
	router.HandleFunc("GET /movies/{id}", getMovieByIDHandler)
	router.HandleFunc("POST /movies", createMovieHandler)
	router.HandleFunc("PUT /movies/{id}", updateMovieHandler)
	router.HandleFunc("DELETE /movies/{id}", deleteMovieHandler)

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	return loggingMiddleware(authMiddleware(router))
}