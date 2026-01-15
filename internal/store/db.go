package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func InitDB() error {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	var err error

	db, err = sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("PostgreSQL database connected")
			return createTables()
		}
		log.Printf("Database not ready yet (Attempt %d/%d). Waiting 2s...\n", i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("could not connect to database after %d attempts: %v", maxRetries, err)
}

func createTables() error {
	queryMovies := `
	CREATE TABLE IF NOT EXISTS movies (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		release_year INTEGER,
		rating REAL, 
		review TEXT
	);`

	if _, err := db.Exec(queryMovies); err != nil {
		return err
	}

	queryGenres := `
    CREATE TABLE IF NOT EXISTS genres (
        id SERIAL PRIMARY KEY,
        name TEXT UNIQUE NOT NULL
    );`

	if _, err := db.Exec(queryGenres); err != nil {
		return err
	}

	queryMovieGenres := `
    CREATE TABLE IF NOT EXISTS movie_genres (
        movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
        genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
        PRIMARY KEY (movie_id, genre_id)
    );`

	if _, err := db.Exec(queryMovieGenres); err != nil {
		return err
	}

	queryInsertGenres := `
    INSERT INTO genres (name) VALUES 
    ('Action'), ('Comédie'), ('Drame'), ('Sci-Fi'), ('Horreur'), ('Aventure')
    ON CONFLICT (name) DO NOTHING;`

	_, err := db.Exec(queryInsertGenres)

	return err
}