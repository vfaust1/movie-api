package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vfaust1/movie-api/internal/store"
)

// --- LE MOCK  ---
type MockMovieStore struct{}

func (m MockMovieStore) GetMovies(title string, filters store.Filters) ([]store.Movie, store.Metadata, error) {
	mockMovies := []store.Movie{
		{ID: 1, Title: "Fake Movie 1", ReleaseYear: 2020},
		{ID: 2, Title: "Fake Movie 2", ReleaseYear: 2021},
	}
	metadata := store.Metadata{TotalRecords: 2, PageSize: 20, CurrentPage: 1}

	return mockMovies, metadata, nil
}

func (m MockMovieStore) AddMovie(movie store.Movie) (store.Movie, error) { return store.Movie{}, nil }
func (m MockMovieStore) GetMoviebyID(id int) (store.Movie, error)        { return store.Movie{}, nil }
func (m MockMovieStore) UpdateMovie(movie store.Movie) error             { return nil }
func (m MockMovieStore) DeleteMovie(id int) error                        { return nil }

// --- LE TEST ---
func TestGetAllMoviesHandler(t *testing.T) {
	mockStore := MockMovieStore{}

	app := &application{
		store: store.Storage{
			Movies: mockStore,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/movies", nil)

	rr := httptest.NewRecorder()

	app.getAllMoviesHandler(rr, req)

	// --- VÉRIFICATIONS ---

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]any
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Impossible de décoder le JSON de réponse")
	}

	moviesList := response["movies"].([]any)
	if len(moviesList) != 2 {
		t.Errorf("On attendait 2 films, on en a reçu %d", len(moviesList))
	}
}
