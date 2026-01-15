package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vfaust1/movie-api/internal/store"
)

func getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	title := queryValues.Get("title")
	page := 1

	if p := queryValues.Get("page"); p!= "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0{
			page = n
		}
	}

	pageSize := 20

	if ps := queryValues.Get("page_size"); ps != "" {
		if n, err := strconv.Atoi(ps); err == nil && n > 0 {
			pageSize = n
		}
			
	}

	sort := "id"
	if s := queryValues.Get("sort"); s != "" {
		sort = s
	}

	filters := store.Filters{
		Page: page,
		PageSize: pageSize,
		Sort: sort,
		SortSafelist: []string{"id", "title", "release_year", "rating"},
	}

	movies, metadata, err := store.GetMovies(title, filters)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error fetching movies :", err)
		return
	}

	response := map[string]any{
		"metadata": metadata,
		"movies": movies,
	}

	respondWithJSON(w, http.StatusOK, response)
}

func getMovieByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be an integer", http.StatusBadRequest)
		return
	}

	movie, err := store.GetMoviebyID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Movie not found", http.StatusNotFound)
		} else {
			log.Println("Error fetching movie:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie store.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := movie.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovie, err := store.AddMovie(movie)

	if err != nil {
		log.Println("Error adding movie:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Movie added: %+v\n", newMovie)

	respondWithJSON(w, http.StatusCreated, newMovie)
}

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be an integer", http.StatusBadRequest)
		return
	}

	err = store.DeleteMovie(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Movie not found", http.StatusNotFound)
		} else {
			log.Println("Error deleting movie:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be an integer", http.StatusBadRequest)
		return
	}

	var movie store.Movie
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	movie.ID = id
	
	if err := movie.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.UpdateMovie(movie)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Movie not found", http.StatusNotFound)
		} else {
			log.Println("Error updating movie:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

// --- HELPERS ---

func respondWithJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}