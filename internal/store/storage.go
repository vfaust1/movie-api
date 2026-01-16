package store

import "database/sql"

// Cette structure contient tous tes modèles
type Storage struct {
	Movies MovieModel
}

// Fonction pour initialiser le Storage avec la connexion DB
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Movies: MovieModel{DB: db},
	}
}