package store

import "database/sql"

type MovieRepository interface {
	AddMovie(Movie) (Movie, error)
	GetMoviebyID(int) (Movie, error)
	GetMovies(string, Filters) ([]Movie, Metadata, error)
	UpdateMovie(Movie) error
	DeleteMovie(int) error
}

type Storage struct {
	Movies MovieRepository
}

// Fonction pour initialiser le Storage avec la connexion DB
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Movies: MovieModel{DB: db},
	}
}