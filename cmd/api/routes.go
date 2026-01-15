package main

import "net/http"

func routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /movies", getAllMoviesHandler)
	router.HandleFunc("GET /movies/{id}", getMovieByIDHandler)
	router.HandleFunc("POST /movies", createMovieHandler)
	router.HandleFunc("PUT /movies/{id}", updateMovieHandler)
	router.HandleFunc("DELETE /movies/{id}", deleteMovieHandler)

	// On emballe le tout dans notre middleware avant de renvoyer
	return loggingMiddleware(authMiddleware(router))
}