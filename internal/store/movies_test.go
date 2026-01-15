package store

import (
	"testing"
	"time"
)

func TestMovie_Validate(t *testing.T) {
	tests := []struct {
		name    string // Nom du test (ex: "Title too short")
		movie   Movie  // La donnée à tester
		wantErr bool   // Est-ce qu'on s'attend à une erreur ?
	}{
		{
			name:    "Valid Movie",
			movie:   Movie{Title: "Inception", ReleaseYear: 2010},
			wantErr: false, // Pas d'erreur
		},
		{
			name:    "Title Too Short",
			movie:   Movie{Title: "A", ReleaseYear: 2010},
			wantErr: true, // Erreur
		},
		{
			name:    "Title Too Long",
			movie:   Movie{Title: "A really really long title that definitely exceeds 50 characters for sure", ReleaseYear: 2010},
			wantErr: true,
		},
		{
			name:    "Year Too Old",
			movie:   Movie{Title: "Old Movie", ReleaseYear: 1800},
			wantErr: true,
		},
		{
			name:    "Future Movie",
			movie:   Movie{Title: "Future", ReleaseYear: time.Now().Year() + 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		// crée un sous-test pour chaque cas
		t.Run(tt.name, func(t *testing.T) {
			err := tt.movie.Validate()

			// Vérification :
			// Si on voulait une erreur mais qu'on a nil -> ÉCHEC
			// Si on ne voulait pas d'erreur mais qu'on en a une -> ÉCHEC
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}