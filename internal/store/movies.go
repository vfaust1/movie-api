package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Rating      *float64 `json:"rating"`
	Review      *string  `json:"review"`
	Genres      []string `json:"genres"`
}

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

// --- FONCTIONS PUBLIQUES (API du package) ---

// Renvoie la liste des films (recherchés si searchTitle
// n'est pas vide ainsi que l'erreur s'il y'en a une.
func GetMovies(searchTitle string, filters Filters) ([]Movie, Metadata, error) {
	orderBy := "id"
	direction := "ASC"

	if filters.Sort != "" {
		if strings.HasPrefix(filters.Sort, "-") {
			direction = "DESC"
			filters.Sort = strings.TrimPrefix(filters.Sort, "-")
		}

		for _, safeValue := range filters.SortSafelist {
			if filters.Sort == safeValue {
				orderBy = safeValue
				break
			}
		}
	}

	limit := filters.PageSize
	offset := (filters.Page - 1) * filters.PageSize

	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, title, release_year, ROUND(rating::numeric, 1), review 
		FROM movies 
		WHERE title ILIKE '%%' || $1 || '%%' 
		ORDER BY %s %s, id ASC
		LIMIT $2 OFFSET $3`, orderBy, direction)

	rows, err := db.Query(query, searchTitle, limit, offset)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	var moviesList []Movie

	for rows.Next() {
		var m Movie

		err := rows.Scan(&totalRecords, &m.ID, &m.Title, &m.ReleaseYear, &m.Rating, &m.Review)
		if err != nil {
			return nil, Metadata{}, err
		}
		moviesList = append(moviesList, m)
	}

	if moviesList == nil {
		moviesList = []Movie{}
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return moviesList, metadata, nil
}

// Ajoute un film et lui attribut un ID,
// renvoie ce même film et nil si l'ajour est bien fait,
// une struct Movie vide et une erreur sinon.
func AddMovie(m Movie) (Movie, error) {
	tx, err := db.Begin()
	if err != nil {
		return Movie{}, err
	}
	defer tx.Rollback()

	queryMovie := `
		INSERT INTO movies (title, release_year, rating, review)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err = tx.QueryRow(
		queryMovie,
		m.Title,
		m.ReleaseYear,
		m.Rating,
		m.Review,
	).Scan(&m.ID)

	if err != nil {
		return Movie{}, err
	}

	for _, genreName := range m.Genres {
		var genreID int
		queryGetGenre := "SELECT id FROM genres WHERE name = $1"

		err = tx.QueryRow(queryGetGenre, genreName).Scan(&genreID)
		if err != nil {
			return Movie{}, fmt.Errorf("genre '%s' not found", genreName)
		}

		queryLink := "INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)"

		_, err = tx.Exec(queryLink, m.ID, genreID)
		if err != nil {
			return Movie{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return Movie{}, err
	}

	return m, nil
}

// Recherche un film par un ID,
// renvoie le film et nil s'il existe,
// une struct Movie vide et une erreur sinon.
func GetMoviebyID(id int) (Movie, error) {
	return getMovieWithGenresSimple(id)
}

func getMovieWithGenresSimple(id int) (Movie, error) {
	queryMovie := `
        SELECT id, title, release_year, ROUND(rating::numeric, 1), review
        FROM movies WHERE id = $1`

	var m Movie
	err := db.QueryRow(queryMovie, id).Scan(&m.ID, &m.Title, &m.ReleaseYear, &m.Rating, &m.Review)
	if err != nil {
		return Movie{}, err
	}

	queryGenres := `
        SELECT g.name FROM genres g
        JOIN movie_genres mg ON g.id = mg.genre_id
        WHERE mg.movie_id = $1`

	rows, err := db.Query(queryGenres, id)
	if err != nil {
		return Movie{}, err
	}
	defer rows.Close()

	var genres []string
	for rows.Next() {
		var g string
		if err := rows.Scan(&g); err != nil {
			continue
		}
		genres = append(genres, g)
	}
	m.Genres = genres

	return m, nil
}

// Supprime un film par son ID,
// renvoie une erreur si elle n'a pas
// pu supprimer ce film, nil sinon.
func DeleteMovie(id int) error {
	query := "DELETE FROM movies WHERE id = $1"

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Met à jour tous les champs d'un Movie,
// renvoie une erreur s'il l'update ne s'est pas fait.
func UpdateMovie(m Movie) error {
	query := `
		UPDATE movies
		SET title = $1, release_year = $2, rating = $3, review = $4
		WHERE id = $5`

	// On execute le query avec les arguments
	res, err := db.Exec(query, m.Title, m.ReleaseYear, m.Rating, m.Review, m.ID)
	if err != nil {
		return err
	}
	// On vérifie que la commande a bien modifié une ligne
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Vérifie que les données du film
// respectent certaines règles métiers
func (m *Movie) Validate() error {
	m.Title = strings.TrimSpace(m.Title)

	// Règle Titre : entre 2 et 50 caractères
	if len(m.Title) < 2 {
		return errors.New("title must be at least 2 characters long")
	}
	if len(m.Title) > 50 {
		return errors.New("title must not exceed 50 characters")
	}

	// Règle Date de sortie : entre 1888 (premier film) et l'année courante
	currentYear := time.Now().Year()
	if m.ReleaseYear < 1888 || m.ReleaseYear > currentYear {
		return errors.New("release_year must be between 1888 and the current year")
	}

	// Règle Note : entre 0 et 10 inclus (mais peut être vide)
	if m.Rating != nil {
		// On met l'étoile *m.Rating pour lire la valeur derrière le pointeur
		if *m.Rating < 0 || *m.Rating > 10 {
			return errors.New("rating must be between 0 and 10")
		}
	}

	// Règle Review : peut être vide mais ne peut exceder 1000 caractères
	if m.Review != nil {
		if len(*m.Review) > 1000 {
			return errors.New("review must not exceed 1000 characters")
		}
	}

	return nil
}

// Permet de calculer les métadonnées
// pour l'affichage
func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	lastPage := (totalRecords + pageSize - 1) / pageSize

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     lastPage,
		TotalRecords: totalRecords,
	}
}
